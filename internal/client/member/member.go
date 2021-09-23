package member

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-gateway/internal/client/eureka"
	"github.com/go-gateway/internal/pkg/maputil"
	"github.com/go-gateway/internal/pkg/seri/member"
	"github.com/spf13/viper"
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
	choosed := maputil.LoadBalance(maputil.MapToString(app.App.Instance, func(instance eureka.Instance) string {
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
	fmt.Printf("b: %s\n", b)
	var r GetMemberResp
	err = member.Decode(b, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}