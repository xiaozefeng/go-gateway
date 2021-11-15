package server

import (
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
)

var ProviderSet = wire.NewSet(NewRouterService)


//type AuthService interface {
//	DetectedService(path string) string
//	FindTarget(serviceId string) string
//	GetReverseProxyPath(path, serviceId string) string
//}
//
//type TokenService interface {
//	CheckToken(token, sourceType string) (memberId int, err error)
//	IsNeedLogin(path, serviceId string) bool
//}

type RouterService struct {
	//AuthService
	//TokenService
	biz.AuthUserCase
	biz.TokenUserCase
}

func NewRouterService(authUserCase biz.AuthUserCase, tokenUserCase biz.TokenUserCase) *RouterService {
	return &RouterService{ authUserCase, tokenUserCase}
}
