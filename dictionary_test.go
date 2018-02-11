package quizzee

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"testing"

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
	ws := readWords(path.Join(getCurrentFileDir(), "dict", "antonym.utf8"))
	for _, w := range ws {
		assert.True(t, IsAntonymWord(w))
	}
	assert.False(t, IsAntonymWord("不行"))
	assert.False(t, IsAntonymWord("用不用"))
	assert.False(t, IsAntonymWord("正确"))
	assert.False(t, IsAntonymWord("正确不正确"))
}

func TestIsIgnoreWord(t *testing.T) {
	ws := readWords(path.Join(getCurrentFileDir(), "dict", "ignore.utf8"))
	for _, w := range ws {
		assert.True(t, IsIgnoreWord(w))
	}
	assert.False(t, IsIgnoreWord("下列的"))
	assert.False(t, IsIgnoreWord("哪部"))
	assert.False(t, IsIgnoreWord("的的"))
}

func TestIsJointWord(t *testing.T) {
	ws := readWords(path.Join(getCurrentFileDir(), "dict", "joint.utf8"))
	for _, w := range ws {
		assert.True(t, IsJointWord(w))
	}
	assert.False(t, IsJointWord("之最"))
}

func TestReplaceWord(t *testing.T) {
	assert.Equal(t, "元素", ReplaceWord("种元素"))
}
