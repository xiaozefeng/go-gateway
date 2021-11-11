package biz

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz/domain"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/util"
)

type AuthRepo interface {
	List() ([]*domain.AuthURL, error)
}

type AuthUserCase struct {
	repo AuthRepo
	cli  *eureka.Client
}

func NewBizUserService(repo AuthRepo, cli *eureka.Client) *AuthUserCase {
	return &AuthUserCase{repo: repo, cli: cli}
}

func (au *AuthUserCase) ListAuthURL() (map[string][]*domain.AuthURL, error) {
	result, err := au.repo.List()
	if err != nil {
		return nil, err
	}
	return convert(result)
}

func (au *AuthUserCase) DetectedService(path string) string {
	return detectService(path)
}

func (au *AuthUserCase) FindTarget(serviceId string) string {
	return findTarget(au.cli, serviceId)
}

func (au *AuthUserCase) GetReverseProxyPath(path, serviceId string) string {
	return getReverseProxyPath(path, serviceId)
}

func findTarget(cli *eureka.Client, serviceId string) string {
	if serviceId == "" {
		return ""
	}
	app, err := cli.GetApp(strings.ToUpper(serviceId))
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
	if len(s) < 2 {
		return path
	}
	return s[1]
}

func getReverseProxyPath(path, serviceId string) string {
	return path[strings.Index(path, serviceId)+len(serviceId):]
}

func convert(list []*domain.AuthURL) (map[string][]*domain.AuthURL, error) {
	var result = make(map[string][]*domain.AuthURL)
	for _, au := range list {
		if v, ok := result[strings.Trim(au.ServiceId, " ")]; ok {
			v = append(v, au)
		} else {
			result[strings.Trim(au.ServiceId, " ")] = make([]*domain.AuthURL, 0)
		}
	}
	return result, nil
}
