package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Employee type
type Employee struct {
	ID   int
	Name string
	City string
}

var db *sql.DB
var err error

func dbConn() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	
	if err != nil {
		log.Fatal(err)
	}
}

func getAllEmployees() []Employee {
	row, err := db.Query("SELECT * FROM employee")
	if err != nil {
    log.Fatal(err)
	}
	emp := Employee{}
	employees := []Employee{}
	for row.Next() {
		err := row.Scan(&emp.ID, &emp.Name, &emp.City)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	return employees
}

func insert(name string, city string)  {
	stmt, err := db.Prepare("INSERT INTO employee(Name,City) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(name, city)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The statement affected %d rows\n", affectedRows)

}

func update(id int, name string, city string) {
	stmt, err := db.Prepare("UPDATE employee SET Name=?, City=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(name, city, id)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The statement affected %d rows\n", affectedRows)
}

func delete(id int) {
	stmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The statement affected %d rows\n", affectedRows)
}

func main() {
	dbConn()
	fmt.Println(getAllEmployees())
	// insert("Zizou", "Algeria")
	// update(1, "test", "Egypt")
	// delete(1)
}
