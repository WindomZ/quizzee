package quizzee

import qdb "github.com/WindomZ/quizzee-db"

var dbTable []byte = nil

func RegisterDB(name string, paths ...string) {
	if name == "" {
		name = "quizzee"
	}
	dbTable = []byte(name)
	if err := qdb.Open(dbTable, paths...); err != nil {
		panic(err)
	}
}

func CloseDB() error {
	return qdb.Close()
}

func HasDB() bool {
	return dbTable != nil
}
