package wire

import (
	"database/sql"

	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api"
	"github.com/xiaozefeng/go-gateway/internal/gateway/api/svc"
	"github.com/xiaozefeng/go-gateway/internal/pkg/middleware"
)

var cache map[string]interface{}

const (
	router_service = "router-service"
	eureka_client  = "eureka-client"
	db_ref         = "db"
)

func InitDI() error {
	eurekaURL := viper.GetString("eureka_url")
	cache = make(map[string]interface{})
	dbRef, err := InitDB("")
	if err != nil {
		return err
	}
	cache[db_ref] = dbRef
	cache[router_service] = InitRouterService(eurekaURL, dbRef)
	cache[eureka_client] = InitEurekaClient(eurekaURL)

	dependyInject()
	return nil
}

func dependyInject() {
	routerService := getRouterService()
	api.SetRouterService(routerService)
	middleware.SetRouterService(routerService)
}

func getRouterService() *svc.RouterService {
	return cache[router_service].(*svc.RouterService)
}

func GetDB() *sql.DB {
	return cache[db_ref].(*sql.DB)
}
