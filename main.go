package main

import (
	"fmt"
	"log"
	"net/http"
	"socialapi/controller"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server is starting...")
	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/signin", controller.Sigin).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// func main() {
// 	fmt.Println("Welcome to Social Api")
// 	var pwd string = "faiz"
// 	byteHash := []byte(pwd)

// 	hash, _ := bcrypt.GenerateFromPassword(byteHash, bcrypt.MinCost)
// 	fmt.Println([]byte(pwd))
// 	fmt.Println(hash)
// 	err := bcrypt.CompareHashAndPassword(hash, byteHash)
// 	if err != nil {
// 		fmt.Println("false")
// 	}
// 	fmt.Println("true")

// }
