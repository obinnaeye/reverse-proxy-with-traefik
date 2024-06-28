package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Service3 runnin on port 8083")
    http.ListenAndServe(":8083", nil)
}
