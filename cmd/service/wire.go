//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"

	"github.com/lechuhuuha/oneshield/pkg/config"
	"github.com/lechuhuuha/oneshield/transports/http"
)

func NewService(cfg *config.Config) (*App, func(), error) {
	panic(wire.Build(
		http.ProviderSet,
		NewApp,
	))
}
