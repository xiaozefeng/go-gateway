package server

import (
	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server/middleware"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var routerSvc *RouterService

func SetRouterService(rs *RouterService) {
	routerSvc = rs
}

func newHandler(mw...gin.HandlerFunc) *gin.Engine {
	g := gin.New()
	var handlers []gin.HandlerFunc
	handlers = append(handlers, gin.Recovery())
	handlers = append(handlers, middleware.NoCache)
	handlers = append(handlers, middleware.Options)
	handlers = append(handlers, middleware.Secure)
	handlers = append(handlers, gin.Logger())
	handlers = append(handlers, mw...)

	g.Use(handlers...)

	g.Any("/*action", func(c *gin.Context) {
		path := c.Request.URL.Path
		logrus.Infof("path: %v", path)

		serviceId := routerSvc.DetectedService(path)
		logrus.Infof("serviceId: %v", serviceId)
		var target = routerSvc.FindTarget(serviceId)
		logrus.Infof("target: %s", target)

		proxy := httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = target
				realPath := routerSvc.GetReverseProxyPath(path, serviceId)
				logrus.Infof("realPath: %v", realPath)
				r.URL.Path = realPath
				logrus.Infof("r.Header: %v", r.Header)
			},
			ErrorHandler: func(rw http.ResponseWriter, r *http.Request, err error) {
				if err != nil {
					c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 500, "resultMsg": "出错了，请一会再试", "data": nil})
				}
			},
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	return g
}

func NewHTTPServer(addr string, handlers ...gin.HandlerFunc) *http.Server {
	gin.SetMode(viper.GetString("runmode"))
	handler := newHandler()
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
