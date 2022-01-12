package http

import (
	"github.com/gorilla/mux"
	"github.com/rajasoun/employee-mgmt/src/employees/service"
)

type EmployeeHandler struct {
	empSvcV1 service.EmployeeService
	*mux.Router
}

func (e EmployeeHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", e.Index).Methods("GET")
	r.HandleFunc("/v1/employees", e.Create).Methods("POST")
}

func NewHandler(empSvcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{empSvcV1, mux.NewRouter()}
	h.SetupRoutes(h.Router)
	return h
}
