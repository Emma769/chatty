package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/emma769/chatty/internal/config"
	"github.com/emma769/chatty/internal/handler"
	"github.com/emma769/chatty/internal/middleware"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/internal/server"
	"github.com/emma769/chatty/internal/services/user"
	"github.com/emma769/chatty/pkg/passlib"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("could not load .env file: %v", err)
	}

	cfg := &config.Getter{}
	lg := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	store, err := psql.NewRepository(cfg)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	defer func() {
		if err := store.Close(); err != nil {
			log.Printf("err closing db: %v", err)
		}
	}()

	if err := store.Ping(); err != nil {
		log.Fatalf("could not ping db: %v", err)
	}

	router := chi.NewRouter()

	router.Use(
		middleware.Recover(lg),
		middleware.Logger(lg),
		middleware.EnableCORS(&middleware.CorsOptions{
			AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				"Accept",
				"Content-Type",
				"Authorization",
				"Host",
			},
		}),
	)

	userService := user.NewService(store, passlib.New())

	handlerService := &handler.Service{
		User: userService,
	}

	api := handler.New(handlerService)

	api.Register(router)

	app := server.New(lg, router, cfg)

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
