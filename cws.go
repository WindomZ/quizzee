package quizzee

import "github.com/WindomZ/gcws"

var cws gcws.CWS

func RegisterCWS(names ...string) {
	var err error
	var name string
	if len(names) != 0 {
		name = names[0]
	}
	cws, err = gcws.NewCWS(name)
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
