package service

import (
	"github.com/rajasoun/employee-mgmt/src/employees/repository"
	"github.com/rajasoun/employee-mgmt/src/entities"
)

type v1Svc struct {
	repo repository.EmployeeRepository
}

func (svc v1Svc) Index() ([]entities.Employee, error) {
	return svc.repo.FindAll()
}

func (svc v1Svc) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(newEmployee)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return v1Svc{repo}
}
