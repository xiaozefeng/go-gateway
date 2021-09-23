package member

import (
	"encoding/json"
)

type APIResult struct {
	ResultCode int         `json:"result_code,omitempty"`
	ResultMsg  string      `json:"result_msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Decode(input []byte, v interface{})error{
	var result APIResult
	err := json.Unmarshal(input, &result)
	if err != nil {
		return  err
	}
	b, err := json.Marshal(result.Data)
	if err != nil {
		return  err
	}
	return  json.Unmarshal(b, &v)
}