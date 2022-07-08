package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func envOrDefault(key, dflt string) string {
	val, ok := os.LookupEnv(key) // Type: string

	if !ok {
		log.Info("using default value:", dflt, " for:", key)
		return dflt
	}

	return val
}

var (
	port     string
	dbDriver string
	dbConn   string
)

func init() {
	port = envOrDefault("PORT", "8000")
	dbDriver = envOrDefault("DB_DRIVER", "postgres")
	dbConn = envOrDefault("DB_CONN", "postgres://localhost:5432/emp-demo?sslmode=disable")
}

func main() {
	var empRepo = repository.NewSQL(dbDriver, dbConn)
	// var empRepo = repository.NewInMem()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port:", port, "...")
	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))

	log.Fatalln(err)
}
