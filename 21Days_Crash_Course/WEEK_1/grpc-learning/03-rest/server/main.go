package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var users = map[string]map[string]string{
	"1": {"id": "1", "name": "John Doe", "city": "New York"},
	"2": {"id": "2", "name": "Jane Smith", "city": "Los Angeles"},
}

func main() {

	http.HandleFunc("/users/", handleUsers)

	fmt.Println("REST server is running on port 8080")

	http.ListenAndServe(":8080", nil)

}

func handleUsers(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {

	case http.MethodGet:
		if id == "" {
			json.NewEncoder(w).Encode(users)
			return
		}

		user, ok := users[id]
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)

	case http.MethodPost:
		body := make(map[string]string)
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		newID := fmt.Sprintf("%d", len(users)+1)
		body["id"] = newID
		users[newID] = body
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(body)

	case http.MethodDelete:
		if _, ok := users[id]; !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		delete(users, id)
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "User with ID %s deleted", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
