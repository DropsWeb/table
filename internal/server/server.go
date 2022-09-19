package server

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/api/createUser", CreateUser)
	http.ListenAndServe("localhost:8181", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}
