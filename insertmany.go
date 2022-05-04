package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	data := []map[string]interface{}{}
	for i := 1; i <= 600; i++ {
		t := strconv.Itoa(i)
		data = append(data, map[string]interface{}{"v1": t, "v2": t + "CHIN YEN", "v3": "LAB ASSISTANT", "v4": "LAB", "v5": time.Now()})
	}
	//	fmt.Println(data)

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:Test@123@tcp(192.168.56.107:3306)/employee")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	stmt, e := db.Prepare("delete from empdata")

	// delete 1st student
	res, e := stmt.Exec()

	// affected rows
	a, e := res.RowsAffected()

	fmt.Println("DELETED ROW, ", a, e)

	sqlStr := "INSERT INTO empdata(EmpID, EmpName, Designation,Department,JoiningDate) VALUES "
	vals := []interface{}{}

	for _, row := range data {
		sqlStr += "(?, ?, ?,?, ?),"
		vals = append(vals, row["v1"], row["v2"], row["v3"], row["v4"], row["v5"])
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	stmt, _ = db.Prepare(sqlStr)

	//format all vals at once
	res, _ = stmt.Exec(vals...)
	fmt.Println(res)

}
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	data := []map[string]interface{}{}
	for i := 1; i <= 600; i++ {
		t := strconv.Itoa(i)
		data = append(data, map[string]interface{}{"v1": t, "v2": t + "CHIN YEN", "v3": "LAB ASSISTANT", "v4": "LAB", "v5": time.Now()})
	}
	//	fmt.Println(data)

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:Test@123@tcp(192.168.56.107:3306)/employee")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	stmt, e := db.Prepare("delete from empdata")

	// delete 1st student
	res, e := stmt.Exec()

	// affected rows
	a, e := res.RowsAffected()

	fmt.Println("DELETED ROW, ", a, e)

	sqlStr := "INSERT INTO empdata(EmpID, EmpName, Designation,Department,JoiningDate) VALUES "
	vals := []interface{}{}

	for _, row := range data {
		sqlStr += "(?, ?, ?,?, ?),"
		vals = append(vals, row["v1"], row["v2"], row["v3"], row["v4"], row["v5"])
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	stmt, _ = db.Prepare(sqlStr)

	//format all vals at once
	res, _ = stmt.Exec(vals...)
	fmt.Println(res)

}
