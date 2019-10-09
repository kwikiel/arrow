package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	var d = (string(requestDump))

	fmt.Fprintf(w, d)
}
