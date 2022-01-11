package repository

import "github.com/rajasoun/employee-mgmt/src/entities"

type EmployeeRepository interface {
	FindAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
