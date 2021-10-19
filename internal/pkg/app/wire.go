//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-gateway/internal/app/gateway/service/auth"
	"github.com/go-gateway/internal/pkg/router/svc"
	"github.com/google/wire"
)
func InitRouterService() *svc.RouterService {
	wire.Build(svc.NewRouterService, auth.NewAuthService, auth.NewTokenService)
	return &svc.RouterService{}
}

