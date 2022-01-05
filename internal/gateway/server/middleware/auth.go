package middleware

import (
	"fmt"
	"github.com/xiaozefeng/go-gateway/internal/gateway/server"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	TOKEN             = "token"
	SourceType        = "sourceType"
	MID               = "mid"
)

var routerSvc *server.RouterService

func Login(c *gin.Context) {
	path := c.Request.URL.Path
	serviceId := routerSvc.DetectedService(path)

	token := c.Request.Header.Get(TOKEN)
	sourceType := c.Request.Header.Get(SourceType)

	memberId, err := routerSvc.CheckToken(token, sourceType)
	log.Infof("memberId: %v", memberId)

	var needLogin = routerSvc.IsNeedLogin(path, serviceId)
	if needLogin && (memberId == -1 || err != nil ) {
			log.Errorf("check token happened err: %v", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 440, "resultMsg": "鉴权失败", "data": nil})
	} else {
		setHeader(c, memberId, sourceType)
		c.Next()
	}
}

func setHeader(c *gin.Context, memberId int, sourceType string) {
	h := c.Request.Header
	h.Set(MID, fmt.Sprintf("%d", memberId))
	h.Set(SourceType, sourceType)
}

func SetRouterService(rs *server.RouterService) {
	routerSvc = rs
}
