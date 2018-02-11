package quizzee

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestAnswer_Parse(t *testing.T) {
	a := NewAnswer("鲁迅：周樟寿")
	assert.NoError(t, a.Parse())
	assert.Equal(t, []string{"鲁迅", "周樟寿"}, a.Keys)
}
