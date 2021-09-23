package biz

import (
	"github.com/go-gateway/internal/data/auth"
	"github.com/go-gateway/internal/data/schema"
)

type AuthService interface {
	ListAuthURL() ([]*schema.AuthURL, error)
}

var authSerice AuthService  = &auth.AuthURLRepo{}

var cache []*schema.AuthURL

func ListAuthURL() ([]*schema.AuthURL, error) {
	if cache != nil {
		return cache, nil
	}
	cache, err := authSerice.ListAuthURL()
	return cache, err
}
