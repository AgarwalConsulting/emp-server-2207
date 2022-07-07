package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/emp-server/entities"
)

func (h *EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	employees, err := h.svcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.svcV1.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(createdEmp)
}
