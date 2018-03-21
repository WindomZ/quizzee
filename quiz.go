package quizzee

import (
	"math"
	"sync"
)

type Quiz struct {
	Question Question
	Answers  Answers
	lock     *sync.Mutex
}

func NewQuiz(question string, answers []string) (qa *Quiz, err error) {
	qa = &Quiz{
		Question: *NewQuestion(question),
		Answers:  *NewAnswers(answers),
		lock:     new(sync.Mutex),
	}
	if err = qa.Question.Parse(); err != nil {
		return
	}
	if err = qa.Answers.Parse(); err != nil {
		return
	}
	return
}

func (qa *Quiz) Recommend() (recommend int, rates []float64) {
	qa.lock.Lock()
	qa.Search(SearchWeightF)
	rates = qa.Answers.Rates()
	if qa.Question.Antonym {
		min := math.MaxFloat64
		for i, rate := range rates {
			if rate <= min {
				min = rate
				recommend = i
			}
		}
	} else {
		var max float64
		for i, rate := range rates {
			if rate >= max {
				max = rate
				recommend = i
			}
		}
	}
	qa.lock.Unlock()
	return
}
