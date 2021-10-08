package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/client/eureka"
	"github.com/go-gateway/internal/pkg/maputil"
	"github.com/spf13/viper"
)

func InitializeRouter(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(mw...)

	g.Any("/*action", func(c *gin.Context) {
		path := c.Request.URL.Path
		log.Infof("path: %v", path)

		serviceId := detectService(path)
		log.Infof("serviceId: %v", serviceId)
		var target = findTarget(serviceId)
		log.Infof("target: %s", target)

		proxy := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = target
				realPath := getReverseProxyPath(path, serviceId)
				log.Infof("realPath: %v", realPath)
				r.URL.Path = realPath
				log.Infof("r.Header: %v", r.Header)
			},
			ErrorHandler: func(rw http.ResponseWriter, r *http.Request, err error) {
				if err != nil {
					c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 500, "resultMsg": "出错了，请一会再试", "data": nil})
				}
			},
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})

}

func findTarget(serviceId string) string {
	var eurekURL = viper.GetString("eureka_url")
	var eurekaClient = eureka.NewClient(eurekURL)
	if serviceId == "" {
		return ""
	}
	app, err := eurekaClient.GetApp(strings.ToUpper(serviceId))
	if err != nil {
		log.Errorf("get service id failed, err: %v", err)
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
