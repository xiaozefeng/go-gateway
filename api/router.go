package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/api/middleware"
	"github.com/go-gateway/internal/client/eureka"
	"github.com/go-gateway/internal/pkg/maputil"
	"github.com/spf13/viper"
)

func InitializeRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Login)
	g.Use(gin.Logger())
	g.Use(mw...)

	g.Any("/*action", func(c *gin.Context) {
		path := c.Request.URL.Path
		log.Printf("path: %v\n", path)

		serviceId := detectService(path)
		log.Printf("service: %v\n", serviceId)
		var target = findTarget(serviceId)
		log.Println("target:", target)

		proxy := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = target
				realPath := getReverseProxyPath(path, serviceId)
				fmt.Printf("realPath: %v\n", realPath)
				r.URL.Path = realPath
				value, exists := c.Get("header")
				if exists {
					var m = value.(map[string]string)
					for k, v := range m {
						r.Header.Set(k, v)
					}
				}

			},
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})

}

func findTarget(serviceId string) string {
	var eurekURL = viper.GetString("eureka_url")
	fmt.Printf("eurekURL: %v\n", eurekURL)
	var eurekaClient = eureka.NewClient(eurekURL)
	if serviceId == "" {
		return ""
	}
	app, err := eurekaClient.GetApp(strings.ToUpper(serviceId))
	if err != nil {
		log.Printf("err: %v\n", err)
		return ""
	}
	return maputil.LoadBalance(maputil.MapToString(app.App.Instance, func(instance eureka.Instance) string {
		if instance.HomePageUrl != "" {
			u, _ := url.Parse(instance.HomePageUrl)
			return u.Host
		}
		return ""
	}))
}

func detectService(path string) string {
	if path == "" {
		return path
	}
	s := strings.Split(path, `/`)
	if len(s) == 0 {
		return path
	}
	return s[1]
}

func getReverseProxyPath(path, serviceId string) string {
	return path[strings.Index(path, serviceId)+len(serviceId):]
}
