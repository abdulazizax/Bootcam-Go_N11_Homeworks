package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func main() {
	client := http.Client{}
	CreateUser(&client)
	GetUsers(&client)
	GetUserById(&client)
	UpdateUserById(&client)
	DeleteUserById(&client)
}

func CreateUser(client *http.Client) {
	user := UserRequest{
		Id:       1,
		Name:     "Komiljon",
		Email:    "komiljon@gmail.com",
		Password: "komil1221",
		Age:      23,
	}

	JsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user: ", err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/user", bytes.NewBuffer(JsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}

	log.Println("CreateUser response status:", resp.Status)
	log.Println("CreateUser response body:", string(body))
	fmt.Printf("-----------------------------------------\n\n")
}

func GetUsers(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8080/users", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}

	log.Println("GetUsers response status:", resp.Status)
	log.Println("GetUsers response body:", string(body))
	fmt.Printf("-----------------------------------------\n\n")
}

func GetUserById(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8080/user?id=1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}

	log.Println("GetUserById response status:", resp.Status)
	log.Println("GetUserById response body:", string(body))
	fmt.Printf("-----------------------------------------\n\n")
}

func UpdateUserById(client *http.Client) {
	user := UserRequest{
		Id:       1,
		Name:     "Salimjon",
		Email:    "salimjon@gmail.com",
		Password: "salim1221",
		Age:      27,
	}

	JsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user: ", err)
		return
	}

	req, err := http.NewRequest("PUT", "http://localhost:8080/user?id=1", bytes.NewBuffer(JsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}

	log.Println("UpdateUserById response status:", resp.Status)
	log.Println("UpdateUserById response body:", string(body))
	fmt.Printf("-----------------------------------------\n\n")
}

func DeleteUserById(client *http.Client) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/user?id=1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return
	}

	log.Println("DeleteUserById response status:", resp.Status)
	log.Println("DeleteUserById response body:", string(body))
	fmt.Printf("-----------------------------------------\n\n")
}
