package quizzee_db

import "regexp"

func TrimQuestion(question string) string {
	if subs := regexp.MustCompile(`^(?:\d*\.)?\s?(.*?)\??？?$`).
		FindStringSubmatch(question); len(subs) > 1 {
		return subs[1]
	}
	return question
}
