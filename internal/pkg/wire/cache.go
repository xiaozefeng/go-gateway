package wire

import (
	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api/svc"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
)

var cache map[string]interface{}

const (
	router_service = "router-service"
	eureka_client  = "eureka-client"
)

func InitDI() {
	eurekaURL := viper.GetString("eureka_url")
	cache = make(map[string]interface{})
	cache[router_service] = InitRouterService(eurekaURL)
	cache[eureka_client] = InitEurekaClient(eurekaURL)
}

func GetRouterService() *svc.RouterService {
	return cache[router_service].(*svc.RouterService)
}

func GetEurekaClient() *eureka.Client {
	return cache[eureka_client].(*eureka.Client)
}
