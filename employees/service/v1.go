package service

import (
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

type empSvcV1 struct {
	repo repository.EmployeeRepository
}

func (svc empSvcV1) Index() ([]entities.Employee, error) {
	return svc.repo.ListAll()
}

func (svc empSvcV1) Create(newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Create(newEmp)
}

// Dependency Injection
func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return &empSvcV1{repo}
}
