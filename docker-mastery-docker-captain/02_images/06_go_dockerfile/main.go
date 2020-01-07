package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from go!!")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":80", nil)
}
