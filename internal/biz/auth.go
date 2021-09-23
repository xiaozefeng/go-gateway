package biz

import "github.com/go-gateway/internal/data/schema"

type AuthService interface {
	ListAuthURL() ([]*schema.AuthURL,error)
}

var authSerice  AuthService

func ListAuthURL() ([]*schema.AuthURL ,error){
	return authSerice.ListAuthURL()
}




