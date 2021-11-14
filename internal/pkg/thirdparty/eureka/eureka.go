package eureka

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/xiaozefeng/go-gateway/internal/pkg/util/httputil"
)

type EurekaServerURL string
type Client struct {
	URL string
}

func (c *Client) GetApps() (*GetAppsResp, error) {
	url := c.URL + "/" + "apps"
	b, err := httputil.Get(url, nil)
	var res GetAppsResp
	err = json.Unmarshal(b, &res)
	return &res, err
}

func (c *Client) GetApp(appId string) (*GetAppResp, error) {
	if len(appId) == 0 {
		return nil, errors.New("appId must not null")
	}
	url := c.URL + "/apps/" + strings.ToUpper(appId)
	b, err := httputil.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var res GetAppResp
	err = json.Unmarshal(b, &res)
	return &res, err
}

func NewClient(url EurekaServerURL) *Client {
	return &Client{URL: string(url)}
}
