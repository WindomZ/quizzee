package quizzee

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
)

type Dictionary struct {
	Words map[string]string
}

func (m Dictionary) Contain(word string) (ok bool) {
	_, ok = m.Words[word]
	return
}

func (m Dictionary) Mapping(word string) string {
	if v, ok := m.Words[word]; ok {
		return v
	}
	return word
}

func NewDictionary(filePath string) (m *Dictionary) {
	m = &Dictionary{
		Words: make(map[string]string),
	}

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}

	reader := bufio.NewReader(f)
	var k, v string
	for {
		size, _ := fmt.Fscanln(reader, &k, &v)
		if size == 0 {
			break
		}
		m.Words[k] = v
	}

	return
}

var (
	antonymDict *Dictionary
	ignoreDict  *Dictionary
	jointDict   *Dictionary
	amendDict   *Dictionary
)

func getCurrentFileDir() string {
	_, filePath, _, _ := runtime.Caller(0)
	return path.Dir(filePath)
}

func init() {
	fileDir := getCurrentFileDir()
	antonymDict = NewDictionary(path.Join(fileDir, "dict", "antonym.utf8"))
	ignoreDict = NewDictionary(path.Join(fileDir, "dict", "ignore.utf8"))
	jointDict = NewDictionary(path.Join(fileDir, "dict", "joint.utf8"))
	amendDict = NewDictionary(path.Join(fileDir, "dict", "amend.utf8"))
}

func IsAntonymWord(w string) bool {
	return antonymDict.Contain(w)
}

func IsIgnoreWord(w string) bool {
	return ignoreDict.Contain(w)
}

func IsJointWord(w string) bool {
	return jointDict.Contain(w)
}

func ReplaceWord(s string) string {
	return amendDict.Mapping(s)
}
