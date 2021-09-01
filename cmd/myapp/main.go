package main

import (
	"fmt"
	"log"
	"main1/pkg"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(pkg.Greet())

	router := mux.NewRouter()
	router.HandleFunc("/", pkg.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", pkg.GetBook).Methods("GET")
	router.HandleFunc("/", pkg.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", pkg.DeleteBook).Methods("DELETE")
	router.HandleFunc("/{id}", pkg.UpdateBook).Methods("PATCH")
	fmt.Println("running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
