package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type logger interface {
	Info(string, ...any)
}

type getter interface {
	GetInt(string, int) int
}

type App struct {
	lg     logger
	router http.Handler
	cfg    getter
}

func New(logger logger, router http.Handler, cfg getter) *App {
	return &App{
		logger,
		router,
		cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.cfg.GetInt("PORT", 8000)),
		ReadTimeout:  time.Duration(a.cfg.GetInt("READ_TIMEOUT", 15)) * time.Second,
		WriteTimeout: time.Duration(a.cfg.GetInt("WRITE_TIMEOUT", 15)) * time.Second,
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
	a.lg.Info("server is starting", "port", strings.TrimPrefix(s.Addr, ":"))
	ch <- s.ListenAndServe()
}

func (a *App) shutdown(s *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.lg.Info("server is shutting down")

	err := s.Shutdown(ctx)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
