package tests

import (
	"testing"

	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/testify/assert"
)

func TestQA_Recommend(t *testing.T) {
	qa, err := quizzee.NewQuiz(
		"手机生产商诺基亚最初是以生产什么为主？",
		[]string{"耳机", "纸", "杂货"},
	)
	assert.NoError(t, err)

	recommend, rates := qa.Recommend()
	assert.True(t, recommend == 1)
	for _, rate := range rates {
		assert.True(t, rate >= 0)
	}
}
