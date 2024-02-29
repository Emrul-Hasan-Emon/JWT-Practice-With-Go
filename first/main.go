package main

import (
	"fmt"
	"net/http"

	"github.com/Emrul-Hasan-Emon/firstJwt/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", handler.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("Coult not start the server. Error: ", err)
	}
}
