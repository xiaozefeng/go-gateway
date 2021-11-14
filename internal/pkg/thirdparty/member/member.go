package member

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member/decode"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member/model"
	"github.com/xiaozefeng/go-gateway/internal/pkg/util"
)

const AppId = "hotel-operation-platform-member"

type UserCase struct {
	cli *eureka.Client
}

func NewUserCase(cli *eureka.Client) *UserCase {
	return &UserCase{cli: cli}
}

func (m *UserCase) GetMember(token, sourceType string) (*model.GetMemberResp, error) {
	app, err := m.cli.GetApp(AppId)
	if err != nil {
		return nil, err
	}
	chosen := util.LoadBalance(util.MapToString(app.App.Instance, func(instance eureka.Instance) string {
		return instance.HomePageUrl
	}))
	var getMemberReq = model.GetMemberReq{
		SourceType: sourceType,
		Token:      token,
	}
	req, _ := json.Marshal(getMemberReq)
	resp, err := http.Post(chosen+"auth/getMember", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Infof("get member result: %s", b)
	var r model.GetMemberResp
	err = decode.Decode(b, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
