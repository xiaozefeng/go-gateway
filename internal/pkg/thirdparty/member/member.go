package member

import (
	"bytes"
	"encoding/json"
	"github.com/google/wire"
	"github.com/xiaozefeng/go-gateway/internal/gateway/biz"
	"github.com/xiaozefeng/go-gateway/internal/pkg/util/mapping"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/eureka"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member/decode"
	"github.com/xiaozefeng/go-gateway/internal/pkg/thirdparty/member/model"
)
var ProviderSet = wire.NewSet(NewUserCase)

const appId = "hotel-operation-platform-member"

type UserCase struct {
	cli *eureka.Client
}

func NewUserCase(cli *eureka.Client) biz.MemberService {
	return &UserCase{cli: cli}
}

func (m *UserCase) GetMember(token, sourceType string) (*model.GetMemberResp, error) {
	app, err := m.cli.GetApp(appId)
	if err != nil {
		return nil, err
	}
	chosen := loadBalance(mapping.MapToString(app.App.Instance, func(instance eureka.Instance) string {
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

func loadBalance(instances []string) string {
	if len(instances) > 0 {
		return instances[0]
	}
	return ""
}
