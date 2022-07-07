package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"algogrit.com/emp-server/entities"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"

	"github.com/gorilla/mux"
)

var empRepo = repository.NewInMem()
var empSvcV1 = service.NewV1(empRepo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	employees, err := empSvcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empSvcV1.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(createdEmp)
}

func main() {
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")
	// r.HandleFunc("/employees", EmployeesHandler)

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8000", r)

	log.Fatalln(err)
}
