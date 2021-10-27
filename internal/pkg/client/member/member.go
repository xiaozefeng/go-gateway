package member

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/xiaozefeng/go-gateway/internal/pkg/client/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member/decode"
	"github.com/xiaozefeng/go-gateway/internal/pkg/client/member/model"
	"github.com/xiaozefeng/go-gateway/internal/pkg/util"
)

const MEMBER_APP_ID = "hotel-operation-platform-member"

type MemberService struct {
	cli *eureka.Client
}

func NewMemberService(cli *eureka.Client) *MemberService {
	return &MemberService{cli: cli}
}

func (m *MemberService) GetMember(token, sourceType string) (*model.GetMemberResp, error) {
	app, err := m.cli.GetApp(MEMBER_APP_ID)
	if err != nil {
		return nil, err
	}
	choosed := util.LoadBalance(util.MapToString(app.App.Instance, func(instance eureka.Instance) string {
		return instance.HomePageUrl
	}))
	var getMemberReq = model.GetMemberReq{
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
	var r model.GetMemberResp
	err = decode.Decode(b, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
