package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (HelloHandler *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprint(w, "hello 70pice")
	} else {
		fmt.Fprintf(w, "hello 70pice!")
	}
}
func main() {
	hello := new(HelloHandler)
	http.ListenAndServe(":9999", hello)
}
