package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"time"

	"SimpleChat/resources"
	"SimpleChat/secrets"
	"github.com/dgrijalva/jwt-go"
)


func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user resources.User
	err := decoder.Decode(&user)
	if err != nil {
		errString := "Error decoding post body"
		log.Printf(err.Error())
		http.Error(w, errString, 500)
		return
	}
	
	err = resources.VerifyUser(&user)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "invalid user name or password", 400)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
    	"name": user.Name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secrets.JwtSecret))
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}