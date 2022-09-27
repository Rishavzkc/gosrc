package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	routing := mux.NewRouter()
	routing.HandleFunc("/employee", CreateEmployee).Methods("POST")
	routing.HandleFunc("/employees", GetEmployees).Methods("GET")
	routing.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
	routing.HandleFunc("/employee/{eid}", UpdateEmployee).Methods("PUT")
	routing.HandleFunc("/employee/{eid}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", routing))
}
