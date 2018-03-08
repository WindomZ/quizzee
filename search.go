package quizzee

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
)

const (
	SearchEngineBaidu = "baidu"
	SearchEngineBing  = "bing"
	SearchEngineSogou = "sogou"
	SearchEngine360   = "360"
)

var SearchWeightF = func(s string) float64 {
	switch s {
	case SearchEngineBaidu:
		return weightBaidu
	case SearchEngineBing:
		return weightBing
	case SearchEngineSogou:
		return weightSogou
	case SearchEngine360:
		return weight360
	}
	return 0
}

func (Quiz) request(url string) (body []byte, err error) {
	c := &http.Client{
		Timeout: Timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36")

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return
}

func (qa *Quiz) searchBaidu(keyword string, f func(string) float64) (err error) {
	factor := f(SearchEngineBaidu)
	if factor <= 0 {
		return
	}

	data, err := qa.request("https://www.baidu.com/s?ie=utf-8&wd=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) searchBing(keyword string, f func(string) float64) (err error) {
	factor := f(SearchEngineBing)
	if factor <= 0 {
		return
	}

	data, err := qa.request("https://cn.bing.com/search?q=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) searchSogou(keyword string, f func(string) float64) (err error) {
	factor := f(SearchEngineSogou)
	if factor <= 0 {
		return
	}

	data, err := qa.request("https://www.sogou.com/web?ie=utf-8&query=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) search360(keyword string, f func(string) float64) (err error) {
	factor := f(SearchEngine360)
	if factor <= 0 {
		return
	}

	data, err := qa.request("https://www.so.com/s?ie=utf-8&q=" +
		url.QueryEscape(keyword))
	if err != nil {
		return
	}
	qa.Answers.Score(string(data), factor)
	return
}

func (qa *Quiz) Search(f func(string) float64) {
	keyword := qa.Question.Keyword()
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		if err := qa.searchBaidu(keyword, f); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.searchBing(keyword, f); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.searchSogou(keyword, f); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	go func() {
		if err := qa.search360(keyword, f); err != nil {
			os.Stderr.WriteString(err.Error())
		}
		wg.Done()
	}()
	wg.Wait()
}
