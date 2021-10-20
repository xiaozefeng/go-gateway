package router

import (
	"net/http"
	"net/http/httputil"

	"github.com/go-gateway/internal/pkg/app"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var svc  = app.GetRouterService()

func Load(g *gin.Engine, mw ...gin.HandlerFunc ) {
	g.Use(mw...)
	
	g.Any("/*action", func(c *gin.Context) {
		path := c.Request.URL.Path
		log.Infof("path: %v", path)

		serviceId := svc.DetectedService(path)
		log.Infof("serviceId: %v", serviceId)
		var target = svc.FindTarget(serviceId)
		log.Infof("target: %s", target)

		proxy := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = target
				realPath := svc.GetReverseProxyPath(path, serviceId)
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

