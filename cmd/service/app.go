package main

import (
	"context"

	"github.com/rs/zerolog/log"

	anyerrgroup "github.com/lechuhuuha/oneshield/pkg/any_errgroup"
	"github.com/lechuhuuha/oneshield/transports/http"
)

type App struct {
	httpService *http.Server
}

func NewApp(
	hv *http.Server,
) *App {
	return &App{
		httpService: hv,
	}
}

func (a *App) Run() error {
	g := anyerrgroup.New()

	defer a.Stop()

	g.Go(func(ctx context.Context) error {
		return a.httpService.Start()
	})

	return g.Wait()
}

func (a *App) Stop() {
	if err := a.httpService.Stop(); err != nil {
		log.Error().Err(err).Msg("failed to stop http server")
	}
}
