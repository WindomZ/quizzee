package quizzee

import "github.com/WindomZ/gcws"

var cws gcws.CWS

func RegisterCWS(name string, paths ...string) {
	var err error
	cws, err = gcws.NewCWS(name, paths...)
	if err != nil {
		panic(err)
	}
	cws.SetConfig(gcws.Config{
		Mode:            gcws.ModeDefault,
		FilterStopWords: true,
	})
}

func HasCWS() bool {
	return cws != nil
}
