package main

import (
    "fmt"
    "net/http"
    "github.com/bitterfly/kuho/server"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	server := &server.Server{}

    http.ListenAndServe(":8080", server)
}