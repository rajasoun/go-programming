package service

import "github.com/rajasoun/employee-mgmt/src/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmployee entities.Employee) (*entities.Employee, error)
}
