package quizzee

import (
	"regexp"
	"strings"
)

type Question struct {
	Text        string   `json:"text"`
	MainText    string   `json:"main_text"`
	CroppedText string   `json:"cropped_text"`
	Words       []string `json:"words"`
	Keys        []string `json:"keys"`
	Antonym     bool     `json:"antonym"`
}

func NewQuestion(text string) *Question {
	return &Question{
		Text: text,
	}
}

var (
	mainTextComp     = regexp.MustCompile(`^([0-9]{1,2}\.)?(.*)?[\?？]$`)
	croppedTextComp1 = regexp.MustCompile(`“(.*)”`)
	croppedTextComp2 = regexp.MustCompile(`(?:(不)?是因为)(.*)`)
)

func (q *Question) cropped() {
	if texts := mainTextComp.
		FindStringSubmatch(q.Text); len(texts) >= 2 {
		q.MainText = strings.TrimSpace(texts[2])
	} else {
		q.MainText = q.Text
	}

	if idx := strings.LastIndex(q.MainText, "，"); idx > 0 {
		q.CroppedText = q.MainText[idx+3:]
	} else {
		q.CroppedText = q.MainText
	}

	if texts := croppedTextComp1.
		FindStringSubmatch(q.CroppedText); len(texts) >= 1 {
		q.CroppedText = strings.TrimSpace(texts[1])
	}

	if texts := croppedTextComp2.
		FindStringSubmatch(q.CroppedText); len(texts) >= 2 {
		if texts[1] != "" {
			q.Antonym = true
		}
		q.CroppedText = strings.TrimSpace(texts[2])
	}
}

func (q *Question) inversion() {
	if count := strings.Count(q.MainText, "不"); count > 0 {
		q.Antonym = count%2 == 0
	}
}

func (q *Question) word2key() {
	var jointWord string
	for _, w := range q.Words {
		if IsIgnoreWord(w) {
			continue
		}
		if IsAntonymWord(w) {
			q.Antonym = true
			continue
		}
		if IsJointWord(w) {
			jointWord = w
			continue
		}
		if jointWord != "" {
			q.Keys = append(q.Keys, jointWord+w)
			jointWord = ""
			continue
		}
		w = ReplaceWord(w)
		for _, k := range q.Keys {
			if k == w {
				w = ""
				break
			}
		}
		if w != "" {
			q.Keys = append(q.Keys, w)
		}
	}
}

func (q *Question) Parse() error {
	q.cropped()
	q.inversion()
	q.Words = cws.Tokenize(q.CroppedText)
	q.Keys = make([]string, 0, len(q.Words))
	q.word2key()

	return nil
}

func (q Question) Keyword() string {
	return strings.Join(q.Keys, " ")
}
