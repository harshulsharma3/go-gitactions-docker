package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	http.HandleFunc("/", getUsers)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{1, "harshul", "harshul@gmail.com"},
		{2, "hitesh", "hitesh@gmail.com"},
		{3, "devops", "devops@gmail.com"},
	}

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
