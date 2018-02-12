package tests

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/testify/assert"
)

func readWords(filePath string) (words []string) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}

	reader := bufio.NewReader(f)
	var w string
	for {
		size, _ := fmt.Fscanln(reader, &w)
		if size == 0 {
			break
		}
		words = append(words, w)
	}
	return
}

func TestIsAntonymWord(t *testing.T) {
	ws := readWords("../dict/antonym.utf8")
	for _, w := range ws {
		assert.True(t, quizzee.IsAntonymWord(w))
	}
	assert.False(t, quizzee.IsAntonymWord("不行"))
	assert.False(t, quizzee.IsAntonymWord("用不用"))
	assert.False(t, quizzee.IsAntonymWord("正确"))
	assert.False(t, quizzee.IsAntonymWord("正确不正确"))
}

func TestIsIgnoreWord(t *testing.T) {
	ws := readWords("../dict/ignore.utf8")
	for _, w := range ws {
		assert.True(t, quizzee.IsIgnoreWord(w))
	}
	assert.False(t, quizzee.IsIgnoreWord("下列的"))
	assert.False(t, quizzee.IsIgnoreWord("哪部"))
	assert.False(t, quizzee.IsIgnoreWord("的的"))
}

func TestIsJointWord(t *testing.T) {
	ws := readWords("../dict/joint.utf8")
	for _, w := range ws {
		assert.True(t, quizzee.IsJointWord(w))
	}
	assert.False(t, quizzee.IsJointWord("之最"))
}

func TestReplaceWord(t *testing.T) {
	assert.Equal(t, "元素", quizzee.ReplaceWord("种元素"))
}
