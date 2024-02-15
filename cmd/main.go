package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/divyatambat/FarmersBasket/internal/api"
	"github.com/divyatambat/FarmersBasket/internal/app"

	"github.com/divyatambat/FarmersBasket/internal/pkg/port"
	repository "github.com/divyatambat/FarmersBasket/internal/repository"

	"github.com/oklog/run"
)

func main() {
	ctx := context.Background()

	// Database connection
	db, err := repository.InitializeDatabase()
	if err != nil {
		//logger.Fatal(err)
	}
	defer db.Close()

	// Services
	services, err := app.NewServices(db)

	// HTTP Router
	router := api.NewRouter(*services)

	// HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port.HTTPPort),
		Handler: router,
	}

	// Run group for concurrent execution and graceful shutdown
	var group run.Group
	group.Add(
		func() error {
			return srv.ListenAndServe()
		},
		func(err error) {
			fmt.Println("Shutting down HTTP server")
			ctx, cancel := context.WithTimeout(ctx, time.Second*30)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				fmt.Println("HTTP server shutdown error:", err)
			} else {
				fmt.Println("HTTP server shutdown complete.")
			}
		},
	)

	// handling for graceful shutdown
	group.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	// Run the group and handle errors
	if err := group.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1) // Exit with a non-zero code to signal failure
	}
}
