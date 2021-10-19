package decode

import (
	"encoding/json"
	"errors"
)

type APIResult struct {
	ResultCode int         `json:"code,omitempty"`
	ResultMsg  string      `json:"msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Decode(input []byte, v interface{}) error {
	var result APIResult
	err := json.Unmarshal(input, &result)
	if err != nil {
		return err
	}
	if result.ResultCode != 100 {
		return errors.New("get member error, message :" + result.ResultMsg)
	}
	b, err := json.Marshal(result.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}
