package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := runApp()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runApp() error {
	if len(os.Args) < 4 {
		return errors.New("app <dbUrl> <httpPort> <staticDir>")
	}

	dbUrl := os.Args[1]
	httpPort := os.Args[2]
	staticDir := os.Args[3]

	err := checkStaticDir(staticDir)
	if err != nil {
		return fmt.Errorf("check static dir: %w", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	err = run(ctx, dbUrl, httpPort, staticDir)
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}

func run(ctx context.Context, dbUrl, httpPort, staticDir string) error {
	dbpool, err := initPostgres(ctx, dbUrl)
	if err != nil {
		return fmt.Errorf("init postgres: %w", err)
	}
	defer dbpool.Close()

	http.Handle("/db", http.HandlerFunc(dbHandler(dbpool)))
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	server := &http.Server{
		Addr: ":" + httpPort,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and server http: %w", err)
		}

		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		err := server.Shutdown(timeoutCtx)
		if err != nil {
			return fmt.Errorf("shutdown http: %w", err)
		}

		return nil
	})

	go func() {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		<-timeoutCtx.Done()

		fmt.Printf("listen on: %s\n", server.Addr)
	}()

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("wait: %w", err)
	}

	return nil
}

func initPostgres(ctx context.Context, url string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("pgxpool new: %w", err)
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("pgxpool ping: %w", err)
	}

	return dbpool, nil
}

func dbHandler(dbpool *pgxpool.Pool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		conn, err := dbpool.Acquire(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Release()

		var ts time.Time

		err = conn.QueryRow(r.Context(), "select now()").Scan(&ts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, ts)
	})
}

func checkStaticDir(staticDir string) error {
	f, err := os.Open(staticDir)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	if !fi.IsDir() {
		return errors.New("static dir is not dir")
	}

	return nil
}
