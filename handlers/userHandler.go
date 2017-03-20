package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"SimpleChat/resources"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	users, err := resources.GetAllUsers()
	if err != nil {
		errString := "Error retrieving collection"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return 
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		errString := "Error transfering data to json"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return 
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user resources.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		errString := "Error decoding post body"
		log.Printf(err.Error())
		http.Error(w, errString, 500)
		return
	}
	
	err = resources.CreateUser(&user)

	if err != nil {
		errString := "Error creating data"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return 
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		errString := "Error transfering data to json"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return 
	}
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := resources.GetUserByID(id)
	if err != nil {
		switch {
			case err == mgo.ErrNotFound:
				errString := "Resource not found"
				log.Printf(errString)
				http.Error(w, errString, 404)
				return
			default:
				errString := "Error transfering data to json"
				log.Print(errString)
				http.Error(w, errString, 500)
				return
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		errString := "Error transfering data to json"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := resources.DeleteUserByID(id)

	switch err {
		case nil:
			w.WriteHeader(http.StatusNoContent)
		case mgo.ErrNotFound:
			log.Printf("Could not find resource with id %s", id)
			w.WriteHeader(http.StatusNoContent)
		default:
			errString := fmt.Sprintf("Error deleting collection %s", id)
			log.Printf(errString)
			http.Error(w, errString, 500)
	}
}

func UserDeleteAll(w http.ResponseWriter, r *http.Request) {
	err := resources.DeleteAllUser()

	switch err {
		case nil:
			w.WriteHeader(http.StatusNoContent)
		default:
			errString := fmt.Sprintf("Error deleting all user collection")
			log.Printf(errString)
			http.Error(w, errString, 500)
	}
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user resources.User
	err := decoder.Decode(&user)
	if err != nil {
		errString := "Error decoding post body"
		log.Printf(errString)
		http.Error(w, errString, 500)
		return
	}
	
	err = resources.VerifyUser(&user)
	if err != nil {
		log.Printf("user does not exist")
		http.Error(w, "user does not exist", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}