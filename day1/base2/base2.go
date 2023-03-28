package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "hello")
	}
}

func hello(w http.ResponseWriter, r http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprint(w, "index")
	}
}
func main() {
	http.ListenAndServe(":9999", nil)
}
