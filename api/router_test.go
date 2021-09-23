package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func Test_str(t *testing.T) {
	path := "/api-server/hello"
	s := strings.Split(path, `/`)
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("s: %v\n", s)
	x := s[1]
	fmt.Printf("x: %v\n", x)
	i := strings.Index(path, x)
	fmt.Printf("i: %v\n", i)
	fmt.Println(path[i+len(x):])

}

func Test_url(t *testing.T) {
	u, _ := url.Parse("http://172.16.32.8:8888/")
	fmt.Printf("u.Host: %v\n", u.Host)
	fmt.Printf("u.Port(): %v\n", u.Port())

}

type resp struct {
	ResultCode int      `json:"resultCode,omitempty"`
	ResultMsg  string      `json:"resultMsg,omitempty"`
	Data       interface {}`json:"data"`
	// GetResult GetResult `json:"data"`
}

type GetResult struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}



func Test_json(t *testing.T) {
	var str = `{"resultCode":100,"resultMsg":"请求成功","data":{"name":"橘子","avatar":"https://img.otmchina.cn/avatar/202109/16/960b9bfc6c2849bc9bde536e96e62641.jpg","mobile":"13751947630"}}`
	var r resp
	err := json.Unmarshal([]byte(str), &r)
	if err != nil {
		t.Error(err)
		return
	}
	b, err := json.Marshal(r.Data)
	if err != nil {
		t.Error(r)
	}
	var getresult GetResult
	err = json.Unmarshal(b, &getresult)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("getresult: %+v\n", getresult)


}
