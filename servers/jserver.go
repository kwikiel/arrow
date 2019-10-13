package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type test_struct struct {
	Test string
	Xd   string
}

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GOT GET ")
	}

	if r.Method == "POST" {
		r.ParseForm()
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)

		if err != nil {
			panic(err)
		}

		fmt.Println(t)
	}
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
