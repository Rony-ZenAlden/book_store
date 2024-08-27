package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// Replace with your MySQL connection details
	username := "root"
	password := "rony_123"
	host := "127.0.0.1"
	port := 3306
	dbname := "book"

	// Create a connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	d, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql"
// )
// func Connect() {
// 	// Replace with your MySQL connection details
// 	username := "root"
// 	password := "rony_123"
// 	host := "127.0.0.1"
// 	port := 3306
// 	dbname := "book"
// 	// Create a connection string
// 	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
// 	// Create a new database session
// 	db, err := sql.Open("mysql", connString)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()
// 	// Ping the database to test the connection
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Connected to MySQL!")
// }
