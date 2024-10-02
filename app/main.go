package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var count int

func main() {

	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	log.Printf("%v : Starting application ", now.Format("2006-01-02 15:04:05"))

	count = 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Router Working")
		fmt.Fprintf(w, "\ntesting")
		log.Printf("API hit %d times", count)
		count++
	})

	fmt.Println("Starting server on :8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8082") // Default listens on :8082
}
