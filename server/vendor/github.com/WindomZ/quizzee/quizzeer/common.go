package quizzeer

import (
	_ "github.com/WindomZ/gcws/sego"
	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/quizzee-db"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

func Register(name string, paths ...string) {
	quizzee.RegisterDB(name, paths...)
	if len(paths) >= 2 {
		quizzee.RegisterCWS("sego", paths[1])
	} else {
		quizzee.RegisterCWS("sego")
	}
}

func Close() error {
	return quizzee.CloseDB()
}

func Recommend(question string, answers []string) (recommend int, rates []float64) {
	return quizzee.Recommend(quizzee_db.TrimQuestion(question), answers)
}

func Mark(question string, answers []string, answer string) error {
	return quizzee.Mark(question, answers, answer)
}
