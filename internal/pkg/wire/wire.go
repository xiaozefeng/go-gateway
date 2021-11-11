//go:build wireinject
// +build wireinject

package wire


import (
	"database/sql"

	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	data "github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/web"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member"
)

func InitRouterService(eurekaServerURL string, db *sql.DB) *web.RouterService {
	panic(wire.Build(web.NewRouterService,
		wire.Bind(new(web.AuthService), new(*biz.AuthUserCase)),
		wire.Bind(new(web.TokenService), new(*biz.TokenService)),
		biz.NewTokenService,
		biz.NewBizUserService,
		member.NewMemberService,
		InitEurekaClient,
		wire.Bind(new(biz.AuthService), new(*biz.AuthUserCase)),
		wire.Bind(new(biz.AuthRepo), new(*data.AuthURLRepo)),
		wire.Bind(new(biz.MemberService), new(*member.MemberService)),
		data.NewAuthURLRepo))
}

func InitEurekaClient(eurekaServerURL string) *eureka.Client {
	panic(wire.Build(eureka.NewClient))
}

func InitDB(url string) (*sql.DB, error) {
	panic(wire.Build(db.Init))
}
