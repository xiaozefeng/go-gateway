//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api/svc"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	data "github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/service/auth"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member"
)

func InitRouterService(eurekaServerURL string, db *sql.DB) *svc.RouterService {
	panic(wire.Build(svc.NewRouterService,
		wire.Bind(new(svc.AuthService), new(*auth.AuthService)),
		wire.Bind(new(svc.TokenService), new(*auth.TokenService)),
		auth.NewAuthService,
		auth.NewTokenService,
		biz.NewBizUserService,
		member.NewMemberService,
		InitEurekaClient,
		wire.Bind(new(auth.BizAuthService), new(*biz.AuthUsercase)),
		wire.Bind(new(biz.AuthRepo), new(*data.AuthURLRepo)),
		wire.Bind(new(auth.MemberService), new(*member.MemberService)),
		data.NewAuthURLRepo))
}

func InitEurekaClient(eurekaServerURL string) *eureka.Client {
	panic(wire.Build(eureka.NewClient))
}

func InitDB() (*sql.DB, error) {
	panic(wire.Build(db.Init))
}
