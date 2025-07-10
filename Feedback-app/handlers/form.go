package handlers

import (
	"net/http"
)

func FeedbackFormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/feedback.html")
}
