package biz

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz/domain"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member/model"
	"strings"
)

type MemberService interface {
	GetMember(token string, sourceType string) (*model.GetMemberResp, error)
}

type TokenUserCase struct {
	authUc *AuthUserCase
	cli    *eureka.Client
	memberSvc MemberService
}

func NewTokenUserCase(authUc *AuthUserCase, cli *eureka.Client, memberSvc MemberService) *TokenUserCase {
	return &TokenUserCase{authUc: authUc, cli: cli, memberSvc: memberSvc}
}

func (ts *TokenUserCase) CheckToken(token, sourceType string) (memberId int, err error) {
	if token == "" || token == "null" || token == "undefined" {
		return -1, fmt.Errorf("invalid token: %s", token)
	}
	resp, err := ts.memberSvc.GetMember(token, sourceType)
	if err != nil {
		return -1, err
	}
	return resp.MemberId, nil
}
func (ts *TokenUserCase) IsNeedLogin(path, serviceId string) bool {
	m, err := ts.authUc.ListAuthURL()
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
		var matchedPath = au.Url == getReverseProxyPath(path, serviceId)
		return serviceIdIsEq && matchedPath
	})

	if len(filtered) == 0 {
		return false
	}
	return filtered[0].IsForeLogin()
}

func filterAuthURL(list []*domain.AuthURL, filter func(*domain.AuthURL) bool) []*domain.AuthURL {
	var res []*domain.AuthURL
	for _, l := range list {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}
