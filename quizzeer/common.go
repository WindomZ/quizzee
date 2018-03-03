package quizzeer

import (
	_ "github.com/WindomZ/gcws/sego"
	"github.com/WindomZ/quizzee"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

func init() {
	quizzee.RegisterCWS("sego")
}

func RegisterDB(name string, paths ...string) {
	quizzee.RegisterDB(name, paths...)
}

func CloseDB() error {
	return quizzee.CloseDB()
}

func Recommend(question string, answers []string) (recommend int, rates []float64) {
	return quizzee.Recommend(question, answers)
}

func Mark(question string, answers []string, answer string) error {
	return quizzee.Mark(question, answers, answer)
}
