//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	data "github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
)

func InitRouterService(eurekaServerURL eureka.ServerURL, mysqlConnectURL db.MySQLConnectURL) (*server.RouterService, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			biz.ProviderSet,
			data.ProviderSet,
			thirdparty.ProviderSet,
			))
}

//func InitDB(url string) (*sql.DB,func(), error) {
//	panic(wire.Build(db.New))
//}
