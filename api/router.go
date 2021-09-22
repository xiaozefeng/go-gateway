package api

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/api/middleware"
)

func InitializeRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect API route.")
	})

	g.Any("/", func(c *gin.Context) {
		path := c.Request.URL.Path

		var service = detectService(path)
		log.Printf("service: %v\n", service)
		var target = findTarget(service)
		log.Println("target:", target)

		proxy := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = target
				r.URL.Path = getReverseProxyPath(path, service)
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
	return "172.16.0.11:9100"
}
func detectService(path string) string {
	return "myapp"
}

func getReverseProxyPath(path, serviceId string) string {
	return "/eureka/apps"
}