//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api/svc"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	data "github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/service/auth"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
)

func InitRouterService(eurekaServerURL string) *svc.RouterService {
	panic(wire.Build(svc.NewRouterService,
		wire.Bind(new(svc.AuthService), new(*auth.AuthService)),
		wire.Bind(new(svc.TokenService), new(*auth.TokenService)),
		auth.NewAuthService,
		auth.NewTokenService,
		biz.NewBizUserService,
		InitEurekaClient,
		wire.Bind(new(auth.BizAuthService), new(*biz.AuthUsercase)),
		wire.Bind(new(biz.AuthRepo), new(*data.AuthURLRepo)),
		data.NewAuthURLRepo))
}

func InitEurekaClient(eurekaServerURL string) *eureka.Client {
	panic(wire.Build(eureka.NewClient))
}
