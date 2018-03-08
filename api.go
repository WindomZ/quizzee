package quizzee

import "encoding/json"

type Response struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	Data    *Request `json:"data"`
}

func (r Response) Bytes() []byte {
	b, _ := json.Marshal(&r)
	return b
}

type Request struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
	Accuracy float64  `json:"accuracy,omitempty"`
}

func (r Request) Success() *Response {
	return &Response{
		Data: &r,
	}
}

func (r Request) Fail(code int, msg string) *Response {
	return &Response{
		ErrCode: code,
		ErrMsg:  msg,
	}
}
