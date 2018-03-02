package tests

import (
	_ "github.com/WindomZ/gcws/jieba"
	"github.com/WindomZ/quizzee"
)

func init() {
	quizzee.RegisterCWS("jieba")
}
