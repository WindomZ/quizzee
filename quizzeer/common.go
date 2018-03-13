package quizzeer

import (
	"path/filepath"
	"runtime"

	_ "github.com/WindomZ/gcws/sego"
	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/quizzee-db"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

func init() {
	_, filePath, _, _ := runtime.Caller(0)
	quizzee.RegisterCWS("sego",
		filepath.Join(filepath.Dir(filepath.Dir(filePath)), "dict", "sego.txt"))
}

func RegisterDB(name string, paths ...string) {
	quizzee.RegisterDB(name, paths...)
}

func CloseDB() error {
	return quizzee.CloseDB()
}

func Recommend(question string, answers []string) (recommend int, rates []float64) {
	return quizzee.Recommend(quizzee_db.TrimQuestion(question), answers)
}

func Mark(question string, answers []string, answer string) error {
	return quizzee.Mark(question, answers, answer)
}
