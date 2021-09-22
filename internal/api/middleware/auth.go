package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	TOKEN             = "token"
	SOURCE_TYPE       = "sourceType"
	SOURCE_TYPE_VALUE = "sourceTypeValye"
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
	// sourceTypeValue:= c.Request.Header.Get(SOURCE_TYPE_VALUE)
	// mid := c.Request.Header.Get(MID)
	// authType:= c.Request.Header.Get(AUTH_TYPE)
	// brandCode:= c.Request.Header.Get(BRAND_CODE)
	// platformId:= c.Request.Header.Get(PLATFORM_ID)

	var needLogin = checkIsNeedLogin(path, serviceId)
	if needLogin {
		memberId, err := checkToken(token, sourceType)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 200, "resultMsg": "登录失败", "data": null})
		}else {
			c.Set(MID, memberId)
		}
	}

}

func checkIsNeedLogin(path, serviceId string) bool {
	
	return false
}

func checkToken(token, sourceType string) (string, error) {
	if token == "" || token == "null" || token == "undefined" {
		return "", fmt.Errorf("invalid token: %s", token)
	}
	return "", nil
}

func getServiceId(path string) string {
	s := strings.Split(path, `/`)
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
