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
	"github.com/spf13/viper"
)

func InitializeRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// g.NoRoute(func(c *gin.Context) {
	// 	c.String(http.StatusNotFound, "the incorrect API route.")
	// })

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
				for k, v := range c.Request.Header {
					if v != nil && len(v) >= 1 {
						r.Header.Set(k, v[0])
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
	return loadBalance(mapToString(app.App.Instance, func(instance eureka.Instance) string {
		if instance.HomePageUrl != "" {
			u, _:= url.Parse(instance.HomePageUrl)
			return u.Host
		}
		return ""
	}))
}

func mapToString(instances []eureka.Instance, apply func(eureka.Instance) string) []string {
	var res []string
	for _, v := range instances {
		res = append(res, apply(v))
	}
	return res
}

func loadBalance(instances []string) string {
	if len(instances) > 0 {
		return instances[0]
	}
	return ""
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
