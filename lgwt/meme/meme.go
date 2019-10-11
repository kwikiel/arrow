package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Url struct {
	Url      string `json:"url"`
	Interval int64  `json:"interval"`
}

func (u Url) String() string {
	return fmt.Sprintf("[%v, %v", u.Url, u.Interval)
}

var db = []Url{}

// curl localhost:8000 -d '{"name":"Hello"}'
func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Url
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db = append(db, Url{msg.Url, msg.Interval})

	for _, kek := range db {
		fmt.Println(kek.Url)
		fmt.Println(kek.Interval)
	}
	
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func main() {
	http.HandleFunc("/", Cleaner)
	address := ":5000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
