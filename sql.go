package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Emp struct {
	empID       int    `json:"id"`
	empName     string `json:"name"`
	designation string `json: "designation"`
	department  string `json: "department"`
	joiningDate string `json: "date"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:Test@123@tcp(192.168.56.107:3306)/employee")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * from empdata")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(results)

	for results.Next() {
		var emp Emp
		// for each row, scan the result into our tag composite object
		err = results.Scan(&emp.empID, &emp.empName, &emp.designation, &emp.department, &emp.joiningDate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		//log.Printf(emp.empID)
		fmt.Println(emp)
	}

}
