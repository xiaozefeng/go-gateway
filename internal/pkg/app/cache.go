package app

import (
	"github.com/go-gateway/internal/pkg/router/svc"
)

var cache map[string]interface{}

func init () {
	cache = make(map[string]interface{})
	cache["router-service"] = InitRouterService()
}

func GetRouterService () *svc.RouterService{
	return cache["router-service"].(*svc.RouterService)
}