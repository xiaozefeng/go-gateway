package httputil

import (
	"io"
	"net/http"
)

func Get(url string, header map[string]string) ([]byte, error) {
	return request(url, http.MethodGet, nil, header)
}

func Post(url string, body []byte, header map[string]string) ([]byte, error) {
	return request(url, http.MethodPost, body, header)
}

func request(url, methond string, body []byte, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	cli := http.DefaultClient
	r, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}
