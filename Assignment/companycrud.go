package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Company struct {
	Id           string `json:"Id"`
	Name         string `json:"Name"`
	Location     string `json:"Location"`
	NoOfEmployee string `json:"NoOfEmpl"`
}

var Companies []Company

func getAllComapnies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Companies)
}

func createCompany(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var company Company
	json.Unmarshal(reqBody, &company)
	Companies = append(Companies, company)
	json.NewEncoder(w).Encode(company)
}

func getCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, company := range Companies {
		if company.Id == key {
			json.NewEncoder(w).Encode(company)
		}
	}
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, company := range Companies {
		if company.Id == id {
			Companies = append(Companies[:index], Companies[index+1:]...)
		}
	}
}

func apiRequests() {

	http.HandleFunc("/companies", getAllComapnies)
	http.HandleFunc("/company", createCompany)
	http.HandleFunc("/company/{id}", getCompanyById)
	http.HandleFunc("/company/{id}", deleteCompany)
	log.Fatal(http.ListenAndServe(":8090", nil))
	fmt.Println("Listing on 8090")
}

func main() {
	Companies = []Company{
		{Id: "1", Name: "Quest", Location: "Pune", NoOfEmployee: "300"},
		{Id: "2", Name: "Tcs", Location: "Delhi", NoOfEmployee: "500"},
		{Id: "3", Name: "Hp", Location: "Banglore", NoOfEmployee: "1000"},
	}
	apiRequests()
}
