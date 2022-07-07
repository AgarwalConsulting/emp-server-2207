package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func main() {
	var empRepo = repository.NewInMem()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	log.Println("Starting server...")
	err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))

	log.Fatalln(err)
}
