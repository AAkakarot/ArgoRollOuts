package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
)

func main() {

    log.Printf("%v : Starting application ", time.Now())
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf(w, "Router Working")
    })

    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
