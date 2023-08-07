package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "shaiq@123"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	defer db.Close()
	var name string
	var password int
	fmt.Println("sign up here")
	fmt.Println("enter name:")
	fmt.Scanln(&name)
	fmt.Println("enter password:")
	fmt.Scanln(&password)
	// create
	insertstmt := `insert into "employee"("name","password") values ($1 , $2)`
	_, e := db.Exec(insertstmt, name, password)
	checkError(e)
	fmt.Println("sign up successfully")

	// Delete
	_, err = db.Exec("DELETE FROM employee WHERE password = 321")
	if err != nil {
		fmt.Println("Failed to delete user:", err)
		return
	}
	//update
	_, err = db.Exec("UPDATE employee SET name = ayesha WHERE id = 27")
	if err != nil {
		return
	}
	fmt.Println("User deleted successfully")
	// read

	rows, err := db.Query("SELECT * FROM employee")
	checkError(err)
	defer rows.Close()

	// Iterate over the rows and print the results
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		checkError(err)
		fmt.Println(name, age)
	}

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
