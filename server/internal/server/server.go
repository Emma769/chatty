package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type logger interface {
	InfoContext(context.Context, string, ...any)
}

type getter interface {
	GetDuration(string) time.Duration
}

type App struct {
	name   string
	port   int
	getter getter
	logger logger
	router http.Handler
}

type AppOptions struct {
	Name   string
	Port   int
	Logger logger
}

func NewWithOptions(
	getter getter,
	router http.Handler,
	opts *AppOptions,
) *App {
	app := &App{
		name:   opts.Name,
		port:   opts.Port,
		logger: opts.Logger,
		router: router,
		getter: getter,
	}

	if app.name == "" {
		app.name = "default name"
	}

	if app.port == 0 {
		app.port = 7000
	}

	if app.logger == nil {
		app.logger = slog.Default()
	}

	return app
}

func (a *App) Run(ctx context.Context) error {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		ReadTimeout:  a.getter.GetDuration("READ_TIMEOUT"),
		WriteTimeout: a.getter.GetDuration("WRITE_TIMEOUT"),
		Handler:      a.router,
	}

	errch := make(chan error)
	go a.listen(s, errch)

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		return a.shutdown(s)
	}
}

func (a *App) listen(s *http.Server, ch chan<- error) {
	a.logger.InfoContext(
		context.Background(),
		"server is starting",
		"port",
		a.port,
	)

	ch <- s.ListenAndServe()
}

func (a *App) shutdown(s *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.logger.InfoContext(context.Background(), "server is shutting down")
	err := s.Shutdown(ctx)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
