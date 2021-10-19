//go:build wireinject
// +build wireinject

package app

import (
	"github.com/go-gateway/internal/app/gateway/biz"
	data "github.com/go-gateway/internal/app/gateway/data/auth"
	"github.com/go-gateway/internal/app/gateway/service/auth"
	"github.com/go-gateway/internal/pkg/router/svc"
	"github.com/google/wire"
)
func InitRouterService() *svc.RouterService {
	wire.Build(svc.NewRouterService, auth.NewAuthService, auth.NewTokenService, biz.NewBizUserService, data.NewAuthURLRepo)
	return &svc.RouterService{}
}

