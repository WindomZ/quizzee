package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "github.com/WindomZ/gcws/jieba"
	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/quizzee-db"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

const (
	size      = 20000       // 计算总次数
	tableName = "questions" // 题库表名
)

func main() {
	quizzee.RegisterCWS("jieba")

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dataPath := filepath.Join(dir, "data", "data.db")
	err = quizzee_db.Open([]byte(tableName), dataPath)
	if err != nil {
		panic(err)
	}
	defer quizzee_db.Close()

	result := map[string]int{
		quizzee.SearchEngineBaidu: 0,
		quizzee.SearchEngineBing:  0,
		quizzee.SearchEngineSogou: 0,
		quizzee.SearchEngine360:   0,
	}

	idx := 0
	sum := quizzee_db.Count()
	err = quizzee_db.Iterator(func(q *quizzee_db.Quiz) bool {
		quiz, err := quizzee.NewQuiz(q.Question, q.Options)
		if err != nil {
			return false
		}

		answer := 0
		for i, opt := range q.Options {
			if opt == q.Answer {
				answer = i
				break
			}
		}

		wg := &sync.WaitGroup{}
		wg.Add(4)

		// Baidu
		go func(quiz quizzee.Quiz) {
			quiz.Search(func(s string) float64 {
				if s == quizzee.SearchEngineBaidu {
					return 1
				}
				return 0
			})
			if maxIndex(quiz.Answers.Scores) == answer {
				result[quizzee.SearchEngineBaidu]++
			}
			wg.Done()
		}(*quiz)

		// Bing
		go func(quiz quizzee.Quiz) {
			quiz.Search(func(s string) float64 {
				if s == quizzee.SearchEngineBing {
					return 1
				}
				return 0
			})
			if maxIndex(quiz.Answers.Scores) == answer {
				result[quizzee.SearchEngineBing]++
			}
			wg.Done()
		}(*quiz)

		// Sogou
		go func(quiz quizzee.Quiz) {
			quiz.Search(func(s string) float64 {
				if s == quizzee.SearchEngineSogou {
					return 1
				}
				return 0
			})
			if maxIndex(quiz.Answers.Scores) == answer {
				result[quizzee.SearchEngineSogou]++
			}
			wg.Done()
		}(*quiz)

		// 360
		go func(quiz quizzee.Quiz) {
			quiz.Search(func(s string) float64 {
				if s == quizzee.SearchEngine360 {
					return 1
				}
				return 0
			})
			if maxIndex(quiz.Answers.Scores) == answer {
				result[quizzee.SearchEngine360]++
			}
			wg.Done()
		}(*quiz)

		time.Sleep(time.Millisecond * time.Duration(500+rand.Int63n(2000)))

		wg.Wait()

		idx++
		fmt.Printf("[%.3f%%] %d/%d\n", float64(idx)/float64(sum), idx, sum)
		return idx <= size
	})
	if err != nil {
		panic(err)
	}

	maxCount := 0
	fmt.Println("------------result------------")
	for name, count := range result {
		if count > maxCount {
			maxCount = count
		}
		fmt.Printf("%6s - %d\n", name, count)
	}
	fmt.Println("------------weight------------")
	for name, count := range result {
		fmt.Printf("%6s - %f\n", name, float64(count)/float64(maxCount))
	}
}

func maxIndex(fs []float64) (ret int) {
	var max float64
	for i, f := range fs {
		if f > max {
			ret = i
		}
	}
	return
}
