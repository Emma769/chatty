package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"

	"github.com/emma769/chatty/internal/config"
	"github.com/emma769/chatty/internal/handler"
	"github.com/emma769/chatty/internal/middleware"
	"github.com/emma769/chatty/internal/server"
)

func main() {
	cfg := &config.Getter{}
	lg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	router := chi.NewRouter()

	router.Use(
		middleware.Recover(lg),
		middleware.Logger(lg),
	)

	api := handler.New(nil)
	api.Register(router)

	app := server.New(lg, router, cfg)

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
