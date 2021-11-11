// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/auth"
	"github.com/xiaozefeng/go-gateway/internal/gateway/data/db"
	"github.com/xiaozefeng/go-gateway/internal/gateway/web"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member"
)

// Injectors from wire.go:

func InitRouterService(eurekaServerURL string, db *sql.DB) *web.RouterService {
	authURLRepo := auth.NewAuthURLRepo(db)
	client := InitEurekaClient(eurekaServerURL)
	authUserCase := biz.NewBizUserService(authURLRepo, client)
	memberService := member.NewMemberService(client)
	tokenService := biz.NewTokenService(authUserCase, client, memberService)
	webRouterService := web.NewRouterService(authUserCase, tokenService)
	return webRouterService
}

func InitEurekaClient(eurekaServerURL string) *eureka.Client {
	client := eureka.NewClient(eurekaServerURL)
	return client
}

func InitDB(url string) (*sql.DB, error) {
	sqlDB, err := db.Init(url)
	if err != nil {
		return nil, err
	}
	return sqlDB, nil
}
