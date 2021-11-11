package wire

import (
	"database/sql"
	"github.com/xiaozefeng/go-gateway/internal/gateway/web"
	"github.com/xiaozefeng/go-gateway/internal/gateway/web/middleware"

	"github.com/spf13/viper"
)

var cache map[string]interface{}

const (
	routerService = "router-service"
	eurekaClient = "eureka-client"
	dbRefKey        = "db"
)

func InitDI() error {
	eurekaURL := viper.GetString("eureka_url")
	cache = make(map[string]interface{})
	dbRef, err := InitDB("")
	if err != nil {
		return err
	}
	cache[dbRefKey] = dbRef
	cache[routerService] = InitRouterService(eurekaURL, dbRef)
	cache[eurekaClient] = InitEurekaClient(eurekaURL)

	dependedInject()
	return nil
}

func dependedInject() {
	routerService := getRouterService()
	web.SetRouterService(routerService)
	middleware.SetRouterService(routerService)
}

func getRouterService() *web.RouterService {
	return cache[routerService].(*web.RouterService)
}

func GetDB() *sql.DB {
	return cache[dbRefKey].(*sql.DB)
}
