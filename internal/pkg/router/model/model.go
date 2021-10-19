package model

type AuthService interface {
	DetectedService(path string) string
	FindTarget(serviceId string) string
	GetReverseProxyPath(path, serviceId string) string
}

type TokenService interface {
	CheckToken(token, sourceType string) (memberId int, err error)
	IsNeedLogin(path ,serviceId string) bool
}

type RouterService struct {
	AuthService
	TokenService
}

func NewRouterService(svc AuthService, tokenService TokenService) *RouterService {
	return &RouterService{svc , tokenService}
}
