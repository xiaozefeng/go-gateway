package wire
//
//import (
//	"github.com/xiaozefeng/go-gateway/internal/gateway/server"
//	"github.com/xiaozefeng/go-gateway/internal/gateway/server/middleware"
//
//	"github.com/spf13/viper"
//)
//
//var cache mapping[string]interface{}
//
//const (
//	routerService = "router-service"
//	eurekaClient = "eureka-thirdparty"
//	dbRefKey        = "db"
//)
//
//func InitDI() error {
//	eurekaURL := viper.GetString("eureka_url")
//	cache = make(mapping[string]interface{})
//	//dbRef, err := InitDB("")
//	//if err != nil {
//	//	return err
//	//}
//	//cache[dbRefKey] = dbRef
//	cache[routerService] = InitRouterService(eurekaURL, dbRef)
//	cache[eurekaClient] = InitEurekaClient(eurekaURL)
//
//	dependedInject()
//	return nil
//}
//
//func dependedInject() {
//	routerService := getRouterService()
//	server.SetRouterService(routerService)
//	middleware.SetRouterService(routerService)
//}
//
//func getRouterService() *server.RouterService {
//	return cache[routerService].(*server.RouterService)
//}
//
////func GetDB() *sql.DB {
////	return cache[dbRefKey].(*sql.DB)
////}
