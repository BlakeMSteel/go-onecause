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
	router := mux.NewRouter()
	router.Use(CORS)
	router.HandleFunc("/", authUser).Methods("POST")
	router.HandleFunc("/", homePage)
	http.Handle("/", router)
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
		log.Printf("Token incorrect\n")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


	if user.Username != User.Username {
		log.Printf("Username incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user.Password != User.Password {
		log.Printf("Password incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
	})
}