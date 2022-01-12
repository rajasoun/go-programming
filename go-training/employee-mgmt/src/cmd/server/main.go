package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/rajasoun/employee-mgmt/src/employees/repository"
	"github.com/rajasoun/employee-mgmt/src/employees/service"

	empHTTP "github.com/rajasoun/employee-mgmt/src/employees/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	msg := "Employee Mgmt !!!"
	fmt.Fprintln(w, msg)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	var empRepo = repository.NewInMem()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	empHandler.SetupRoutes(r)
	log.Println("Starting server on port:", 8000, "...")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
