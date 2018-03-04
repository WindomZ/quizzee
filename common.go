package quizzee

import (
	"fmt"

	qdb "github.com/WindomZ/quizzee-db"
)

func Recommend(question string, answers []string) (recommend int, rates []float64) {
	recommend = -1

	// db
	if HasDB() {
		q := qdb.GetQuiz(question)
		if q.Completion() {
			for i, answer := range answers {
				if answer == q.Answer {
					recommend = i
					break
				}
			}
			if recommend >= 0 {
				rates = make([]float64, len(answers))
				rates[recommend] = 1 // 100%
				return
			}
		}
	}

	// search
	if HasCWS() {
		q, err := NewQuiz(question, answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		recommend, rates = q.Recommend()
		return
	}

	return
}

func Mark(question string, answers []string, answer string) error {
	if question == "" || answer == "" {
		return nil
	}
	// db
	if HasDB() {
		q := qdb.GetQuiz(question)
		if len(answers) != 0 {
			q.Options = answers
		}
		q.Answer = answer
		return q.Store()
	}
	return nil
}
