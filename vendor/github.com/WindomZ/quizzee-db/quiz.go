package quizzee_db

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Quiz struct {
	Question string   `json:"q"`
	Options  []string `json:"os"`
	Answer   string   `json:"a"`
	Update   int64    `json:"ts"`
}

func (q Quiz) Valid() bool {
	return q.Question != ""
}

func (q Quiz) Completion() bool {
	return q.Valid() && q.Answer != ""
}

func (q *Quiz) trim() {
	q.Question = strings.TrimSpace(q.Question)
	for i, opt := range q.Options {
		q.Options[i] = strings.TrimSpace(opt)
	}
	q.Answer = strings.TrimSpace(q.Answer)
}

func (q Quiz) Hash() []byte {
	h := md5.New()
	h.Write([]byte(q.Question))
	return h.Sum(nil)
}

func (q Quiz) Bytes() []byte {
	b, _ := json.Marshal(&q)
	return b
}

func (q *Quiz) Store() error {
	q.trim()
	if !q.Completion() {
		return fmt.Errorf("invalid quiz")
	}
	q.Update = time.Now().Unix()
	return Put(q.Hash(), q.Bytes())
}

func GetQuiz(question string) *Quiz {
	q := new(Quiz)
	if question = strings.TrimSpace(question); question == "" {
		return q
	}

	q.Question = question
	if b := Get(q.Hash()); len(b) != 0 {
		json.Unmarshal(b, q)
	}
	return q
}

func ParseQuiz(data []byte) *Quiz {
	q := new(Quiz)
	json.Unmarshal(data, q)
	if q.Valid() && !q.Completion() {
		if b := Get(q.Hash()); len(b) != 0 {
			json.Unmarshal(b, q)
		}
	}
	return q
}

func NewQuiz(question string) *Quiz {
	return &Quiz{Question: question}
}
