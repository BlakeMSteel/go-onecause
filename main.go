package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type LoginFormat struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

var User LoginFormat

func main() {
	User = LoginFormat{
		Username: "c137@onecause.com",
		Password: "#th@nH@rm#y#r!$100%D0p#",
	}
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/auth", authUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func authUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user = User
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Printf("Body parse error, %v", err)
		w.WriteHeader(400)
		return
	}

	tokenString := r.Header.Get("Authorization")
	hour, minute, _ := time.Now().UTC().Clock()
	correctToken := fmt.Sprintf("%02d%02d", hour, minute)
	if tokenString != correctToken {
		log.Printf("Token incorrect")
		w.WriteHeader(401)
		return
	}


	if user.Username != User.Username {
		log.Printf("Username incorrect")
		w.WriteHeader(401)
		return
	}

	if user.Password != User.Password {
		log.Printf("Password incorrect")
		w.WriteHeader(401)
		return
	}

	w.WriteHeader(200)
}