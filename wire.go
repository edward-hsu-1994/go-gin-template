// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go-fiber-template/accesses"
	"go-fiber-template/routes"
	"go-fiber-template/services"
)

func FilberConfig() ([]fiber.Config, error) {
	return []fiber.Config{}, nil
}

func InitializeApp() (*fiber.App, error) {
	wire.Build(
		FilberConfig,
		fiber.New)

	return nil, nil
}

var repoSet = wire.NewSet(
	accesses.NewMockPostRepository,
)

var serviceSet = wire.NewSet(
	services.NewPostService,
)

var routesSet = wire.NewSet(
	routes.NewNewsRouter,
	routes.NewPostRouter,
	AssembleFiberRouters,
)

func AssembleFiberRouters(
	newsRouter *routes.NewsRouter,
	postRouter *routes.PostRouter,
) []routes.FiberRouter {
	return []routes.FiberRouter{newsRouter, postRouter}
}

func InitialFiberRouters() ([]routes.FiberRouter, error) {
	wire.Build(
		repoSet,
		serviceSet,
		routesSet,
	)
	return nil, nil
}
