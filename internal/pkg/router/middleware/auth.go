package middleware

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/pkg/app"
)

var (
	TOKEN             = "token"
	SOURCE_TYPE       = "sourceType"
	SOURCE_TYPE_VALUE = "sourceTypeValue"
	MID               = "mid"
	AUTH_TYPE         = "authType"
	BRAND_CODE        = "brandCode"
	PLATFORM_ID       = "platformId"
)

var svc = app.GetRouterService()

func Login(c *gin.Context) {
	path := c.Request.URL.Path
	serviceId := svc.DetectedService(path)

	token := c.Request.Header.Get(TOKEN)
	sourceType := c.Request.Header.Get(SOURCE_TYPE)

	memberId, err := svc.CheckToken(token, sourceType)
	log.Infof("memberId: %v", memberId)

	var needLogin = svc.IsNeedLogin(path, serviceId)
	if needLogin {
		if err != nil {
			log.Errorf("check token happened err: %v", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 440, "resultMsg": "鉴权失败", "data": nil})
		} else {
			setHeader(c, memberId, sourceType)
			c.Next()
		}
	} else {
		setHeader(c, memberId, sourceType)
		c.Next()
	}
}

func setHeader(c *gin.Context, memberId int, sourceType string) {
	h := c.Request.Header
	h.Set(MID, fmt.Sprintf("%d", memberId))
	h.Set(SOURCE_TYPE, sourceType)
}

