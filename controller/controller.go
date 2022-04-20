package controller

import (
	"encoding/json"
	"net/http"
	"socialapi/helper"
	"socialapi/model"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var usermodel model.UserModel
	_ = json.NewDecoder(r.Body).Decode(&usermodel)
	err := helper.AddUserToDB(usermodel)
	if err == "nil" {
		json.NewEncoder(w).Encode("user sucessfully added")
	} else {
		json.NewEncoder(w).Encode(err)
	}
	// fmt.Println("email:" + usermodel.Email)
	// fmt.Println("pass " + usermodel.Password)
	// fmt.Println(usermodel.ID)

	// passhash, _ := bcrypt.GenerateFromPassword([]byte(usermodel.Password), bcrypt.MinCost)

	// usermodel.Password = string(passhash)
	// fmt.Println(usermodel.Password)

}

func Sigin(w http.ResponseWriter, r *http.Request) {

	var usermodel model.UserModel
	json.NewDecoder(r.Body).Decode(&usermodel)
	if helper.AuthChecker(usermodel) {
		json.NewEncoder(w).Encode("Login Sucessfull")
	} else {
		json.NewEncoder(w).Encode("Invalid credential")
	}

}
