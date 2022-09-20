package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DropsWeb/table/internal/database"
)

func Server() {
	http.HandleFunc("/api/createUser", CreateUser)
	http.ListenAndServe("localhost:8181", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	database.CreateData(user)
}
