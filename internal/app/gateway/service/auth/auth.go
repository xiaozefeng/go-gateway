package auth

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-gateway/internal/app/gateway/biz"
	"github.com/go-gateway/internal/app/gateway/data/schema"
	"github.com/go-gateway/internal/pkg/client/eureka"
	"github.com/go-gateway/internal/pkg/client/member"
	"github.com/go-gateway/internal/pkg/router/svc"
	"github.com/go-gateway/internal/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthService struct {
}

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


type TokenService struct {

}

func NewTokenService() svc.TokenService {
	return &TokenService{}
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
	m, err := biz.ListAuthURL()
	if err != nil {
		logrus.Errorf("list auth url happened error: %v", err)
		return false
	}
	v, ok := m[serviceId]
	if !ok {
		logrus.Errorf("service id :%s can not find any url", serviceId)
		return false
	}
	filtered := filterAuthURL(v, func(au *schema.AuthURL) bool {
		var serviceIdIsEq = strings.ToLower(serviceId) == strings.ToLower(au.ServiceId)
		var mathcedPath = au.Url == getReverseProxyPath(path, serviceId)
		return serviceIdIsEq && mathcedPath
	})

	if len(filtered) == 0 {
		return false
	}
	first := filtered[0]
	return first.ForceAuth == 1
}

func filterAuthURL(list []*schema.AuthURL, filter func(*schema.AuthURL) bool) []*schema.AuthURL {
	var res []*schema.AuthURL
	for _, l := range list {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}
