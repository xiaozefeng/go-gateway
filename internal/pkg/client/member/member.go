package member

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member/decode"
	"github.com/xiaozefeng/go-gateway/internal/pkg/util"
)

type GetMemberResp struct {
	MemberId   int    `json:"memberId,omitempty"`
	SourceType string `json:"sourceType,omitempty"`
}

type GetMemberReq struct {
	SourceType string `json:"sourceType"`
	Token      string `json:"token"`
}

var MEMBER_APP_ID = "hotel-operation-platform-member"

func GetMember(token, sourceType string) (*GetMemberResp, error) {
	cli := eureka.NewClient(viper.GetString("eureka_url"))
	app, err := cli.GetApp(MEMBER_APP_ID)
	if err != nil {
		return nil, err
	}
	choosed := util.LoadBalance(util.MapToString(app.App.Instance, func(instance eureka.Instance) string {
		return instance.HomePageUrl
	}))
	var getMemberReq = GetMemberReq{
		SourceType: sourceType,
		Token:      token,
	}
	req, _ := json.Marshal(getMemberReq)
	resp, err := http.Post(choosed+"auth/getMember", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Infof("get member result: %s", b)
	var r GetMemberResp
	err = decode.Decode(b, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
