//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-gateway/internal/app/gateway/service/auth"
	"github.com/go-gateway/internal/pkg/router/model"
	"github.com/google/wire"
)
func InitRouterService() *model.RouterService {
	wire.Build(model.NewRouterService, auth.NewAuthService, auth.NewTokenService)
	return &model.RouterService{}
}

