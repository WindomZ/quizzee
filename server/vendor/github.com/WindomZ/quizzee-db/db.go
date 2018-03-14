package quizzee_db

import (
	"encoding/json"

	"github.com/WindomZ/gkv"
)

func Open(table []byte, paths ...string) error {
	return gkv.Open(table, paths...)
}

func Close() error {
	return gkv.Close()
}

func Put(key, value []byte) error {
	return gkv.Put(key, value)
}

func Get(key []byte) []byte {
	return gkv.Get(key)
}

func Count() int {
	return gkv.Count()
}

func Iterator(f func(*Quiz) bool) error {
	return gkv.Iterator(func(k []byte, v []byte) bool {
		if q := new(Quiz); json.Unmarshal(v, q) == nil {
			return f(q)
		}
		return true
	})
}
