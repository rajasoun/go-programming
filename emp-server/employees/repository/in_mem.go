package repository

import (
	"sync"

	"algogrit.com/httpex/entities"
)

type inMemRepo struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inMemRepo) FindAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()
	return repo.employees, nil
}

func (repo *inMemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()
	newEmp.ID = len(repo.employees) + 1
	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Pranav", "Cloud", 1002},
		{3, "Raja", "SRE", 10001},
	}

	return &inMemRepo{employees: employees}
}
