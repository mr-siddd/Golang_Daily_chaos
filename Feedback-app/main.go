package main

import (
	"Feedback-app/handlers" // Update this import path to match your module name, e.g., "github.com/yourusername/Feedback-app/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/feedback", handlers.FeedbackFormHandler)
	http.HandleFunc("/submit", handlers.SubmitHandler)
	http.HandleFunc("/feedbacks", handlers.FeedbackListHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
