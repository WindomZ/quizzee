package quizzee

import "strings"

type Answer struct {
	Text        string   `json:"text"`
	CroppedText string   `json:"cropped_text"`
	Words       []string `json:"words"`
	Keys        []string `json:"keys"`
}

func NewAnswer(text string) *Answer {
	return &Answer{
		Text: text,
	}
}

func (a *Answer) Parse() error {
	a.CroppedText = strings.TrimSpace(strings.Replace(a.Text,
		"：", "", 1))
	a.Words = cws.Tokenize(a.CroppedText)
	a.Keys = make([]string, 0, len(a.Words))

	for _, w := range a.Words {
		if len(w) > 1 {
			a.Keys = append(a.Keys, w)
		}
	}

	return nil
}

func (a Answer) Score(s string) (count float64) {
	count = float64(strings.Count(s, a.CroppedText))
	for _, k := range a.Keys {
		count += 0.3 * float64(strings.Count(s, k))
	}
	return
}

type Answers struct {
	Answers []*Answer
	Scores  []float64
}

func NewAnswers(texts []string) *Answers {
	a := &Answers{
		Answers: make([]*Answer, 0, len(texts)),
	}
	for _, text := range texts {
		if text != "" {
			a.Answers = append(a.Answers, NewAnswer(text))
		}
	}
	a.Scores = make([]float64, len(a.Answers))
	return a
}

func (a Answers) Size() int {
	return len(a.Answers)
}

func (a *Answers) Parse() (err error) {
	for _, ans := range a.Answers {
		if err = ans.Parse(); err != nil {
			return
		}
	}
	return
}

func (a *Answers) Score(s string, factor float64) {
	if s != "" && factor > 0 {
		for i, ans := range a.Answers {
			a.Scores[i] += factor * ans.Score(s)
		}
	}
}

func (a *Answers) Rates() []float64 {
	var sum float64
	for _, score := range a.Scores {
		sum += score
	}
	rate := make([]float64, len(a.Scores))
	for i, score := range a.Scores {
		rate[i] = score / sum
	}
	return rate
}
