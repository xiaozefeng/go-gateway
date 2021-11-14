//go:build wireinject
//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	data "github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member"
)

func InitRouterService(eurekaServerURL eureka.ServerURL, mysqlConnectURL db.MySQLConnectURL) (*server.RouterService,func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			biz.ProviderSet,
			data.ProviderSet,

			wire.Bind(new(server.AuthService), new(*biz.AuthUserCase)),
			wire.Bind(new(server.TokenService), new(*biz.TokenUserCase)),
			wire.Bind(new(biz.AuthRepo), new(*data.URLRepo)),
			wire.Bind(new(biz.MemberService), new(*member.UserCase))))
}

//func InitDB(url string) (*sql.DB,func(), error) {
//	panic(wire.Build(db.New))
//}
