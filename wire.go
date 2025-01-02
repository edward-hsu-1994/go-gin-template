// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-fiber-template/accesses"
	"go-fiber-template/routes"
	"go-fiber-template/services"
)

func GinConfig() ([]gin.OptionFunc, error) {
	return []gin.OptionFunc{}, nil
}

func InitializeApp() (*gin.Engine, error) {
	wire.Build(
		GinConfig,
		gin.Default)

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
	AssembleGinRouters,
)

func AssembleGinRouters(
	newsRouter *routes.NewsRouter,
	postRouter *routes.PostRouter,
) []routes.GinRouter {
	return []routes.GinRouter{newsRouter, postRouter}
}

func InitialGinRouters() ([]routes.GinRouter, error) {
	wire.Build(
		repoSet,
		serviceSet,
		routesSet,
	)
	return nil, nil
}
