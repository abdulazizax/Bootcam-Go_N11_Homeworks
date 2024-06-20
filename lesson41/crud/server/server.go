package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

var Users []UserRequest

func main() {
	h := http.NewServeMux()

	h.HandleFunc("POST /user", CreateUser)
	h.HandleFunc("GET /user", GetUserById)
	h.HandleFunc("GET /users", GetUsers)
	h.HandleFunc("PUT /user", UpdateUserById)
	h.HandleFunc("DELETE /user", DeleteUserById)

	log.Println("Server is working on :8080 port ...")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Users = append(Users, user)
	w.Write([]byte("User created!"))
	w.WriteHeader(http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if len(Users) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(Users)
	w.WriteHeader(http.StatusOK)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	for _, v := range Users {
		if v.Id == intID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var user UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, v := range Users {
		if v.Id == intID {
			Users[i].Name = user.Name
			Users[i].Email = user.Email
			Users[i].Password = user.Password
			Users[i].Age = user.Age
			w.Write([]byte("User updated!"))
			w.WriteHeader(http.StatusOK)
			log.Println(Users)
			log.Println(v)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	for i, v := range Users {
		if v.Id == intID {
			Users = append(Users[:i], Users[i+1:]...)
			w.Write([]byte("User deleted!"))
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
