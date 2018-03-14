package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/quizzee/quizzeer"
)

const port = 8080

func handlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp quizzee.Response
	w.Write(resp.Bytes())
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req quizzee.Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.Write(req.Fail(10, err.Error()).Bytes())
		return
	}
	defer r.Body.Close()

	if req.Question == "" {
		w.Write(req.Fail(20, "empty question").Bytes())
		return
	}
	if len(req.Options) == 0 {
		w.Write(req.Fail(21, "empty options").Bytes())
		return
	}

	if req.Answer != "" {
		err := quizzeer.Mark(req.Question, req.Options, req.Answer)
		if err != nil {
			w.Write(req.Fail(30, err.Error()).Bytes())
			return
		}
		req.Accuracy = 1 // 100%
	} else {
		recommend, rates := quizzeer.Recommend(req.Question, req.Options)
		req.Answer = req.Options[recommend]
		req.Accuracy = rates[recommend]
	}

	w.Write(req.Success().Bytes())
}

func main() {
	var tableName, dataPath string
	flag.StringVar(&tableName, "t", "quizzee", "the table name of database")
	flag.StringVar(&dataPath, "f", "quizzee.db", "the path of database file")

	flag.Parse()

	fmt.Printf("load database: '%s' in '%s'\n", tableName, dataPath)

	quizzeer.Register(tableName, dataPath, "sego.txt")
	defer quizzeer.Close()

	fmt.Println("listen port:", port)

	http.HandleFunc("/ping", handlerPing)
	http.HandleFunc("/", handlerRoot)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
