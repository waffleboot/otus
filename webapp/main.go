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
	"github.com/waffleboot/app/port/repo"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	return runWithContext(ctx)
}

func runWithContext(ctx context.Context) error {
	var (
		mode      string
		portNum   int
		connStr   string
		staticDir string
	)

	rootFlagSet := flag.NewFlagSet("", flag.ContinueOnError)
	rootFlagSet.IntVar(&portNum, "port", 0, "http port")
	rootFlagSet.StringVar(&staticDir, "static", "", "static dir")
	rootFlagSet.StringVar(&mode, "mode", "", "webapp mode")

	err := rootFlagSet.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	storage, err := files.NewStorage(staticDir)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	var metadata repo.Metadata

	switch mode {
	case "memory":
		metadata = database.NewMemoryDatabase()
	case "postgres":
		pg, err := database.NewPostgresClient(ctx, connStr, 5*time.Second)
		if err != nil {
			return fmt.Errorf("new postgres client: %w", err)
		}
		defer pg.Close()

		err = pg.Initialize(ctx)
		if err != nil {
			return fmt.Errorf("initialize: %w", err)
		}

		metadata = pg
	default:
		return fmt.Errorf("unsupported mode: %s", mode)
	}

	svc := domain.NewService(metadata, storage)

	err = startServer(ctx, svc, portNum)
	if err != nil {
		return fmt.Errorf("start server: %w", err)
	}

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
