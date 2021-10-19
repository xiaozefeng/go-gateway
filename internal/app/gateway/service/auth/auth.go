package auth

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-gateway/internal/app/gateway/biz/domain"
	"github.com/go-gateway/internal/pkg/client/eureka"
	"github.com/go-gateway/internal/pkg/client/member"
	"github.com/go-gateway/internal/pkg/router/svc"
	"github.com/go-gateway/internal/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthService struct{}

func NewAuthService() svc.AuthService {
	return &AuthService{}
}

func (as *AuthService) DetectedService(path string) string {
	return detectService(path)
}

func (as *AuthService) FindTarget(serviceId string) string {
	return findTarget(serviceId)
}

func (as *AuthService) GetReverseProxyPath(path, serviceId string) string {
	return getReverseProxyPath(path, serviceId)
}

func findTarget(serviceId string) string {
	var eurekURL = viper.GetString("eureka_url")
	var eurekaClient = eureka.NewClient(eurekURL)
	if serviceId == "" {
		return ""
	}
	app, err := eurekaClient.GetApp(strings.ToUpper(serviceId))
	if err != nil {
		logrus.Errorf("get service id failed, err: %v", err)
		return ""
	}
	return util.LoadBalance(util.MapToString(app.App.Instance, func(instance eureka.Instance) string {
		if instance.HomePageUrl != "" {
			u, _ := url.Parse(instance.HomePageUrl)
			return u.Host
		}
		return ""
	}))
}

func detectService(path string) string {
	if path == "" {
		return path
	}
	s := strings.Split(path, `/`)
	if len(s) == 0 {
		return path
	}
	return s[1]
}

func getReverseProxyPath(path, serviceId string) string {
	return path[strings.Index(path, serviceId)+len(serviceId):]
}

type BizAuthService interface {
	ListAuthURL() (map[string][]*domain.AuthURL, error)
}

type TokenService struct {
	BizAuthService
}

func NewTokenService(ba BizAuthService) svc.TokenService {
	return &TokenService{ba}
}

func (ts *TokenService) CheckToken(token, sourceType string) (memberId int, err error) {
	if token == "" || token == "null" || token == "undefined" {
		return -1, fmt.Errorf("invalid token: %s", token)
	}
	resp, err := member.GetMember(token, sourceType)
	if err != nil {
		return -1, err
	}
	return resp.MemberId, nil
}
func (ts *TokenService) IsNeedLogin(path, serviceId string) bool {
	m, err := ts.ListAuthURL()
	if err != nil {
		logrus.Errorf("list auth url happened error: %v", err)
		return false
	}
	v, ok := m[serviceId]
	if !ok {
		logrus.Errorf("service id :%s can not find any url", serviceId)
		return false
	}
	filtered := filterAuthURL(v, func(au *domain.AuthURL) bool {
		var serviceIdIsEq = strings.ToLower(serviceId) == strings.ToLower(au.ServiceId)
		var mathcedPath = au.Url == getReverseProxyPath(path, serviceId)
		return serviceIdIsEq && mathcedPath
	})

	if len(filtered) == 0 {
		return false
	}
	return filtered[0].IsForeLogin()
}

func filterAuthURL(list []*domain.AuthURL, filter func(*domain.AuthURL) bool) []*domain.AuthURL{
	var res []*domain.AuthURL
	for _, l := range list {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}
