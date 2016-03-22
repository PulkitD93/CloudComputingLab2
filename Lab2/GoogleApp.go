package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(rsp http.ResponseWriter, r *http.Request) {
    fmt.Fprint(rsp, "Hello, This is Lab 2 Google App Engine Application !")
	fmt.Fprint(rsp, "\n");
	}