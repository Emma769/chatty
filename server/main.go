package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"github.com/emma769/chatty/internal/config"
	"github.com/emma769/chatty/internal/handler"
	"github.com/emma769/chatty/internal/repository/psql"
	"github.com/emma769/chatty/internal/server"
	"github.com/emma769/chatty/internal/services/session"
	"github.com/emma769/chatty/internal/services/user"
	"github.com/emma769/chatty/internal/tokens"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("could not load .env file: %v", err)
	}

	getter := new(config.Getter)

	var (
		PORT          = getter.GetInt("PORT")
		POSTGRES_URI  = getter.GetString("POSTGRES_URI")
		SYMMETRIC_KEY = getter.GetString("SYMMETRIC_KEY")
	)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	repo, err := psql.NewRepository(POSTGRES_URI, getter)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	defer func() {
		if err := repo.Close(); err != nil {
			log.Printf("err closing db: %v", err)
		}
	}()

	var (
		user    = user.NewService(repo)
		session = session.NewService(repo)
	)

	maker, err := tokens.NewMaker(SYMMETRIC_KEY)
	if err != nil {
		log.Fatal(err)
	}

	handlerService := &handler.Service{
		User:    user,
		Session: session,
	}

	h := handler.NewWithService(context.Background(), getter, maker, logger, handlerService)

	appOptions := &server.AppOptions{
		Logger: logger,
		Port:   PORT,
	}

	app := server.NewWithOptions(getter, h.ApiRoutes(), appOptions)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
