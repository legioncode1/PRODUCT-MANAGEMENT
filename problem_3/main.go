package main

import (
    "fmt"
    "net/http"
    "time"
    _ "github.com/gorilla/mux" // Add a dummy dependency
)


func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC1123)
    fmt.Fprintf(w, "Current Date and Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", currentTimeHandler)
    fmt.Println("Server is listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

