package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type NodeType struct {
	IsDir bool
	Name  string
	Path  string
}

func main() {

	http.HandleFunc("/file/", func(w http.ResponseWriter, req *http.Request) {
		filePath := "public" + (strings.TrimPrefix(req.URL.Path, "/file"))
		file, err := ioutil.ReadFile(filePath)

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("없는파일"))
		}

		setupResponse(&w, req)
		w.Write(file)
	})

	http.HandleFunc("/dir/", func(w http.ResponseWriter, req *http.Request) {
		dirPath := "public" + (strings.TrimPrefix(req.URL.Path, "/dir"))
		dir, err := ioutil.ReadDir(dirPath)
		var data []NodeType

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("없는 폴더"))
		}

		for _, file := range dir {
			data = append(data, NodeType{IsDir: file.IsDir(), Name: file.Name(), Path: dirPath})
		}

		json, err := json.Marshal(data)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("??"))
		}

		setupResponse(&w, req)
		w.Write([]byte(string(json)))
	})

	http.ListenAndServe(":5000", nil)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
