package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
}

type AppConfig struct {
	db   *sql.DB
	port string
}

func main() {
	// load env file if necessary
	loadEnvFile()

	// initialize database
	db := initDB()
	defer db.Close()

	// get port number
	portNumber := os.Getenv("PORT")

	// create an app
	app := &AppConfig{
		db:   db,
		port: portNumber,
	}

	// start the server
	app.serve()
}

// connect to sqlite
func initDB() *sql.DB {
	conn, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// if .env file exists, load environment variables from the file
// if it doesn't exist and error, ignore error
func loadEnvFile() {
	_ = godotenv.Load(".env")
}

// start http server
func (app *AppConfig) serve() {
	http.HandleFunc("/employees", app.employeesHandler)
	err := http.ListenAndServe(":"+app.port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// get employees from the database
func (app *AppConfig) getEmployees() []Employee {
	employees := []Employee{}
	rows, err := app.db.Query("SELECT * FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		employee := Employee{}
		rows.Scan(&employee.Id, &employee.Gender)
		employees = append(employees, employee)
	}
	return employees
}

// handle the request to get employees
func (app *AppConfig) employeesHandler(w http.ResponseWriter, r *http.Request) {
	employees := app.getEmployees()
	jsonData, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonData)
}
