package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	user     string
	password string
	db       string
}

var dbdata = database{user: "root", password: "12345", db: "db"}

var connectData string = dbdata.user + ":" + dbdata.password + "@tcp(127.0.0.1:3310)/" + dbdata.db

func Db() {
	db, err := sql.Open("mysql", connectData)
	if err != nil {
		log.Fatal(err)
	}
	tables := [2]string{"user", "room"}

	for i := 0; i < 2; i++ {
		_, table_check := db.Query("select * from " + tables[i] + ";")

		if table_check != nil {

			switch tables[i] {
			case "room":
				parameters := []string{
					"id INT PRIMARY KEY AUTO_INCREMENT",
					"name VARCHAR(255)",
					"city VARCHAR(255)",
					"active BOOLEAN DEFAULT true",
				}
				CreateTable(tables[i], parameters)

			case "user":
				parameters := []string{
					"id INT PRIMARY KEY AUTO_INCREMENT",
					"name VARCHAR(255)",
					"city VARCHAR(255)",
					"active BOOLEAN DEFAULT true",
					"phone VARCHAR(255)",
				}
				CreateTable(tables[i], parameters)
			}

		}
	}

}

func CreateTable(table string, columns []string) {
	query := "CREATE TABLE " + table + "("
	for i := 0; i < len(columns); i++ {
		query = query + columns[i]
		if i+1 < len(columns) {
			query = query + ","
		}
	}
	query = query + ")"
	fmt.Println(query)
	db, err := sql.Open("mysql", connectData)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err1 := db.Exec(query)
	if err1 != nil {
		log.Fatal(err, 123)
	} else {
		fmt.Println("database success created")
	}
}

// func CreateData() {
// 	db, err := sql.Open("mysql", connectData)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// func GetData() {

// }
