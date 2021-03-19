package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("lol %s %s\n", r.Method, r.URL.Path)
    fmt.Fprintf(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
