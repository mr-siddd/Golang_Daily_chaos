package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type Feedback struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	name := r.FormValue("name")
	message := r.FormValue("message")

	newFeedback := Feedback{
		Name:    name,
		Message: message,
	}

	// Read existing file or initialize
	data, err := os.ReadFile("feedback.json")
	if err != nil {
		data = []byte("[]")
	}

	var feedbacks []Feedback
	json.Unmarshal(data, &feedbacks)

	// Append new entry
	feedbacks = append(feedbacks, newFeedback)

	// Marshal and save
	finalData, _ := json.MarshalIndent(feedbacks, "", "  ")
	os.WriteFile("feedback.json", finalData, 0644)

	// Redirect to list
	http.Redirect(w, r, "/feedbacks", http.StatusSeeOther)
}
