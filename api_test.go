package quizzee

import (
	"encoding/json"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestResponse_Bytes(t *testing.T) {
	resp := &Response{
		ErrCode: 1,
		ErrMsg:  "error",
	}
	b1 := resp.Bytes()
	b2, _ := json.Marshal(resp)
	assert.Equal(t, b1, b2)
}

func TestRequest_Fail(t *testing.T) {
	req := &Request{}
	resp := req.Fail(1, "error")
	assert.Equal(t, 1, resp.ErrCode)
	assert.Equal(t, "error", resp.ErrMsg)
}

func TestRequest_Success(t *testing.T) {
	req := &Request{}
	resp := req.Success()
	assert.Empty(t, resp.ErrCode)
	assert.Empty(t, resp.ErrMsg)
}
