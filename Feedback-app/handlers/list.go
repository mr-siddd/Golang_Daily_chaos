package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func FeedbackListHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("feedback.json")
	if err != nil {
		http.Error(w, "Unable to read feedback file", http.StatusInternalServerError)
		return
	}

	var feedbacks []Feedback
	json.Unmarshal(data, &feedbacks)

	fmt.Fprintf(w, "<h2>All Feedbacks</h2>")
	for _, fb := range feedbacks {
		fmt.Fprintf(w, "<b>%s:</b> %s<br><br>", fb.Name, fb.Message)
	}
}
