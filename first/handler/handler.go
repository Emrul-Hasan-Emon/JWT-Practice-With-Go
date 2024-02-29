package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Emrul-Hasan-Emon/firstJwt/jwt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	// Fetch the user credentials from the request body
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("The user request body: ", user)

	// For the first time of verification we are checking whether the credentials are matched or not
	if user.Name == "Emon" && user.Password == "123456" {
		tokenString, err := jwt.CreateToken(user.Name) // Trying to generate a JWT Token
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed to Generate Token")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString) // We are sending the JWT token along with the repsonse
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized) // user credentials doesn't get matched
		fmt.Fprint(w, "Invalid Credentials")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization") // Authorization header will contain the token

	fmt.Println("Token String: ", tokenString)

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	tokenString = tokenString[len("Bearer "):] // Removing "Bearer " from the very beginning because we need only the token

	err := jwt.VerifyToken(tokenString) // Verifying the token
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Token")
		return
	}
	fmt.Fprint(w, "Welcome to the protected resources")
}
