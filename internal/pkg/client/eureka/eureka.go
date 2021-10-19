package eureka

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-gateway/internal/pkg/httputil"
)

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

func (c *Client) GetApp(appid string) (*GetAppResp, error) {
	if len(appid) == 0 {
		return nil, errors.New("appid must not null")
	}
	url := c.URL + "/apps/" + strings.ToUpper(appid)
	b, err := httputil.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var res GetAppResp
	err = json.Unmarshal(b, &res)
	return &res, err
}

func NewClient(url string) *Client {
	return &Client{URL: url}
}
