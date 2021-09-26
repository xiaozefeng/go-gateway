package biz

import (
	"github.com/go-gateway/internal/data/auth"
	"github.com/go-gateway/internal/data/schema"
)

type AuthService interface {
	ListAuthURL() ([]*schema.AuthURL, error)
}

var authSerice AuthService = &auth.AuthURLRepo{}

// var cache []*schema.AuthURL

var cache map[string][]*schema.AuthURL

func ListAuthURL() (map[string][]*schema.AuthURL, error) {
	if cache != nil {
		return cache, nil
	}
	list, err := authSerice.ListAuthURL()
	if err != nil {
		return nil, err
	}

	cache, err := convert(list)
	return cache, err
}

func convert(list []*schema.AuthURL) (map[string][]*schema.AuthURL, error) {
	var result = make(map[string][]*schema.AuthURL)
	for _, au := range list {
		auList := result[au.ServiceId]
		auList = append(auList, au)
	}
	return result, nil
}
