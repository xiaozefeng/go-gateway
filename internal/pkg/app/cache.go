package app

import "github.com/go-gateway/internal/pkg/router/model"

var cache map[string]interface{}

func init () {
	cache = make(map[string]interface{})
	cache["router-service"] = InitRouterService()
}

func GetRouterService () *model.RouterService{
	return cache["router-service"].(*model.RouterService)
}