package quizzeer

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestRegister(t *testing.T) {
	Register("testing", "../data/data.db")
}

func TestRecommend(t *testing.T) {
	recommend, rates := Recommend(
		"手机生产商诺基亚最初是以生产什么为主？",
		[]string{"耳机", "纸", "杂货"},
	)
	assert.True(t, recommend == 1)
	for _, rate := range rates {
		assert.True(t, rate >= 0)
	}
}

func TestMark(t *testing.T) {
	assert.NoError(t, Mark(
		"手机生产商诺基亚最初是以生产什么为主？",
		[]string{"耳机", "纸", "杂货"},
		"纸",
	))
}

func TestClose(t *testing.T) {
	assert.NoError(t, Close())
}
