package quizzee

import (
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type Quiz struct {
	Question Question
	Answers  Answers
	lock     sync.Mutex
}

func NewQuiz(question string, answers []string) (qa *Quiz, err error) {
	qa = &Quiz{
		Question: *NewQuestion(question),
		Answers:  *NewAnswers(answers),
	}
	if err = qa.Question.Parse(); err != nil {
		return
	}
	if err = qa.Answers.Parse(); err != nil {
		return
	}
	return
}

const Timeout = time.Second * 3

func (Quiz) request(url string) (body []byte, err error) {
	c := &http.Client{
		Timeout: Timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return
}

func (qa *Quiz) searchBaidu(keyword string, factor float64) (err error) {
	data, err := qa.request("https://www.baidu.com/s?ie=utf-8&wd=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) searchBing(keyword string, factor float64) (err error) {
	data, err := qa.request("https://cn.bing.com/search?q=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) searchSogou(keyword string, factor float64) (err error) {
	data, err := qa.request("https://www.sogou.com/web?ie=utf-8&query=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) search360(keyword string, factor float64) (err error) {
	data, err := qa.request("https://www.so.com/s?ie=utf-8&q=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

const (
	weightBaidu = 0.3
	weightBing  = 0.5
	weightSogou = 0.6
	weight360   = 0.5
)

func (qa *Quiz) search() {
	keyword := qa.Question.Keyword()
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		if err := qa.searchBaidu(keyword, weightBaidu); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.searchBing(keyword, weightBing); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.searchSogou(keyword, weightSogou); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.search360(keyword, weight360); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
}

func (qa *Quiz) Recommend() (recommend int, rates []float64) {
	qa.lock.Lock()
	qa.search()
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
