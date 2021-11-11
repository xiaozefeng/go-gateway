package httputil

import (
	"bytes"
	"io"
	"net/http"
)

func Get(url string, header map[string]string) ([]byte, error) {
	return request(url, http.MethodGet, nil, header)
}

func Post(url string, body []byte, header map[string]string) ([]byte, error) {
	return request(url, http.MethodPost, body, header)
}

func request(url, method string, body []byte, header map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil{
		bodyReader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, bodyReader)
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
