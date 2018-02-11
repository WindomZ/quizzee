package quizzee

import (
	"github.com/WindomZ/gcws"
	_ "github.com/WindomZ/gcws/jieba"
)

var cws gcws.CWS

func init() {
	var err error
	cws, err = gcws.NewCWS("jieba")
	if err != nil {
		panic(err)
	}
	cws.SetConfig(gcws.Config{
		Mode:            gcws.ModeDefault,
		FilterStopWords: true,
	})
}
