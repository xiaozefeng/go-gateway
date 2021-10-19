package biz

import (
	"strings"

	"github.com/go-gateway/internal/app/gateway/biz/domain"
)

type AuthRepo interface {
	List() ([]*domain.AuthURL, error)
}

type AuthUsercase struct {
	AuthRepo
}

func NewBizUserService(repo AuthRepo)  *AuthUsercase{
	return &AuthUsercase{repo}
}

func (au *AuthUsercase) ListAuthURL() (map[string][]*domain.AuthURL, error) {
	result, err := au.List()
	if err != nil {
		return nil, err
	}
	return convert(result)
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
