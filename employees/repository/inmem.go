package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmemRepo struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()
	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()
	newEmployee.ID = len(repo.employees) + 1
	repo.employees = append(repo.employees, newEmployee)

	return &newEmployee, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Bhanu", "Cloud", 10002},
		{3, "Hrishikesh", "SRE", 1003},
	}

	return &inmemRepo{employees: employees}
}
