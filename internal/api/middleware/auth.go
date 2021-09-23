package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-gateway/internal/client/member"
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
			log.Printf("err: %v\n", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"resultCode": 500, "resultMsg": "鉴权失败", "data": nil})
		}else {
			var m = make(map[string]string)
			m[SOURCE_TYPE] = sourceType
			m[MID] = fmt.Sprintf("%d", memberId)
			c.Set("header", m)
			c.Next()
		}
	}else {
		c.Next()
	}
	

}

func checkIsNeedLogin(path, serviceId string) bool {	
	return false
}

func checkToken(token, sourceType string) (int, error) {
	if token == "" || token == "null" || token == "undefined" {
		return -1, fmt.Errorf("invalid token: %s", token)
	}
	resp, err := member.GetMember(token, sourceType)
	if err != nil {
		return -1, err
	}
	log.Printf("resp: %v\n", resp)
	return resp.MemberId, nil
}

func getServiceId(path string) string {
	s := strings.Split(path, `/`)
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
