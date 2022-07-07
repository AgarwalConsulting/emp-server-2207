package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Bhanu", "Cloud", 10002},
	{3, "Hrishikesh", "SRE", 1003},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	log.Println("Employees:", employees)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmployee)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else {
// 		EmployeesIndexHandler(w, req)
// 	}
// }

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
