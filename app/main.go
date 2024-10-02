package main

import "fmt"

// app/main.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf(w, "Hello, Argo CD!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Printf("Failed to start server: %v", err)
    }
}
