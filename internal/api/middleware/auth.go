package middleware

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/app/gateway/biz"
	"github.com/go-gateway/internal/app/gateway/data/schema"
	"github.com/go-gateway/internal/pkg/client/member"
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

func Login(c *gin.Context) {
	path := c.Request.URL.Path
	serviceId := getServiceId(path)

	token := c.Request.Header.Get(TOKEN)
	sourceType := c.Request.Header.Get(SOURCE_TYPE)

	memberId, err := checkToken(token, sourceType)
	log.Infof("memberId: %v", memberId)

	var needLogin = checkIsNeedLogin(path, serviceId)
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

func checkIsNeedLogin(path, serviceId string) bool {
	m, err := biz.ListAuthURL()
	if err != nil {
		log.Errorf("list auth url happened error: %v", err)
		return false
	}
	v, ok := m[serviceId]
	if !ok {
		log.Errorf("service id :%s can not find any url", serviceId)
		return false
	}
	filtered := filterAuthURL(v, func(au *schema.AuthURL) bool {
		var serviceIdIsEq = strings.ToLower(serviceId) == strings.ToLower(au.ServiceId)
		var mathcedPath = au.Url == getReverseProxyPath(path, serviceId)
		return serviceIdIsEq && mathcedPath
	})

	if len(filtered) == 0 {
		return false
	}
	first := filtered[0]
	return first.ForceAuth == 1
}

func getReverseProxyPath(path, serviceId string) string {
	return path[strings.Index(path, serviceId)+len(serviceId):]
}

func filterAuthURL(list []*schema.AuthURL, filter func(*schema.AuthURL) bool) []*schema.AuthURL {
	var res []*schema.AuthURL
	for _, l := range list {
		if filter(l) {
			res = append(res, l)
		}
	}
	return res
}

func checkToken(token, sourceType string) (int, error) {
	if token == "" || token == "null" || token == "undefined" {
		return -1, fmt.Errorf("invalid token: %s", token)
	}
	resp, err := member.GetMember(token, sourceType)
	if err != nil {
		return -1, err
	}
	return resp.MemberId, nil
}

func getServiceId(path string) string {
	s := strings.Split(path, `/`)
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
