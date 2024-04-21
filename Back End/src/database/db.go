package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DB_USERNAME = "lab"
const DB_PASSWORD = "stefan"
const DB_NAME = "catalog"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *sql.DB

func InitDb() *sql.DB {
	Db = connectDB()
	return Db
}

func connectDB() *sql.DB {
	var err error
	ruta := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

	db, err := sql.Open("mysql", ruta)

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection to database : error=%v", err)
		return nil
	}

	return db
}

func CloseDB(db *sql.DB) {
	defer db.Close()
}
