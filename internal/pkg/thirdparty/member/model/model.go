package model

type GetMemberResp struct {
	MemberId   int    `json:"memberId,omitempty"`
	SourceType string `json:"sourceType,omitempty"`
}

type GetMemberReq struct {
	SourceType string `json:"sourceType"`
	Token      string `json:"token"`
}
