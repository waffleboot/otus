package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/waffleboot/app/adapter/database"
	"github.com/waffleboot/app/adapter/files"
	"github.com/waffleboot/app/adapter/web"
	"github.com/waffleboot/app/domain"
	"golang.org/x/sync/errgroup"
)

func main() {

	var (
		portNum   int
		connStr   string
		staticDir string
	)

	flag.IntVar(&portNum, "port", 0, "http port")
	flag.StringVar(&connStr, "conn", "", "conn string")
	flag.StringVar(&staticDir, "static", "", "static dir")
	flag.Parse()

	err := run(portNum, connStr, staticDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(portNum int, connStr, staticDir string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pg, err := database.NewPostgresClient(ctx, connStr, 5*time.Second)
	if err != nil {
		return fmt.Errorf("new postgres client: %w", err)
	}

	err = pg.Initialize(ctx)
	if err != nil {
		return fmt.Errorf("initialize: %w", err)
	}

	storage, err := files.NewStorage(staticDir)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	svc := domain.NewService(pg, storage)

	err = startServer(ctx, svc, portNum)
	if err != nil {
		return fmt.Errorf("start server: %w", err)
	}

	pg.Close()

	return nil
}

func startServer(ctx context.Context, svc *domain.Service, port int) error {

	srv, err := web.NewServer(web.WithHttpPort(port), web.WithService(svc))
	if err != nil {
		return fmt.Errorf("new server: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := srv.Start()
		if err != nil {
			return fmt.Errorf("start server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("shutdown server: %w", err)
		}

		fmt.Println("server stopped")

		return nil
	})

	g.Go(func() error {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		select {
		case <-timeoutCtx.Done():
			fmt.Println("server started")
			return nil
		case <-ctx.Done():
			return nil
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("wait: %w", err)
	}

	return nil
}
