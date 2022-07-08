package repository

import (
	"log"

	"github.com/jmoiron/sqlx"

	"algogrit.com/emp-server/entities"
)

type sqlRepo struct {
	*sqlx.DB
}

func (repo *sqlRepo) ListAll() ([]entities.Employee, error) {
	employees := []entities.Employee{}
	err := repo.DB.Select(&employees, "SELECT * FROM employees ORDER BY name ASC")

	return employees, err
}

func (repo *sqlRepo) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	var empCount int
	rows, err := repo.DB.Query("SELECT count(*) FROM employees")

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&empCount)
	}

	newEmployee.ID = empCount + 1

	_, err = repo.DB.NamedExec("INSERT INTO employees (id, name, department, project_id) VALUES (:id, :name, :department, :project_id)", newEmployee)

	return &newEmployee, err
}

func NewSQL(dbDriver, dbConn string) EmployeeRepository {
	db, err := sqlx.Connect(dbDriver, dbConn)

	if err != nil {
		log.Fatalln("Unable to connect:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id numeric primary key, name text, department text, project_id numeric)")

	if err != nil {
		log.Fatalln("Unable to create table:", err)
	}

	return &sqlRepo{db}
}
