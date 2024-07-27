package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/emma769/chatty/internal/config"
	"github.com/emma769/chatty/internal/server"
)

func main() {
	cfg := &config.Getter{}
	lg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	router := chi.NewRouter()

	app := server.New(lg, router, cfg)

	if err := app.Run(context.TODO()); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
