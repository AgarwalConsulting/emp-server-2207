package repository

import "algogrit.com/emp-server/entities"

type inmemRepo struct {
	employees []entities.Employee
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmemRepo) Create(newEmployee entities.Employee) (*entities.Employee, error) {
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
