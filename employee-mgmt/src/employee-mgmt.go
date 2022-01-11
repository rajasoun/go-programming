package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"division"`
	ProjectID  int    `json:"project"`
}

func (e Employee) String() string {
	return fmt.Sprintf("%s is working in %s department on project: %d.", e.Name, e.Department, e.ProjectID)
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Pranav", "Cloud", 1002},
	{3, "Raja", "SRE", 10001},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmployee)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"
		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe("127.0.0.1:8000", handlers.LoggingHandler(os.Stdout, r))
}
