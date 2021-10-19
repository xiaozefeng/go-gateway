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
	panic(wire.Build(svc.NewRouterService,
		wire.Bind(new(svc.AuthService), new(*auth.AuthService)),
		wire.Bind(new(svc.TokenService), new(*auth.TokenService)),
		auth.NewAuthService,
		auth.NewTokenService,
		biz.NewBizUserService,
		wire.Bind(new(auth.BizAuthService), new(*biz.AuthUsercase)),
		wire.Bind(new(biz.AuthRepo), new(*data.AuthURLRepo)),
		data.NewAuthURLRepo))
}
