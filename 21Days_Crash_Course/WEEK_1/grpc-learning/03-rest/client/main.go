package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	base := "http://localhost:8080"

	// GET all users
	resp, _ := http.Get(base + "/users/")
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET /users/ →", string(body))

	// GET one user
	resp, _ = http.Get(base + "/users/1")
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("GET /users/1 →", string(body))

	// POST new user
	resp, _ = http.Post(
		base+"/users/",
		"application/json",
		strings.NewReader(`{"name":"Arjun","city":"Pune"}`),
	)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("POST /users/ →", string(body))

	// GET all again — should show new user
	resp, _ = http.Get(base + "/users/")
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("GET /users/ after POST →", string(body))
}
