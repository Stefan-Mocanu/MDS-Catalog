package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func InitDb() *sql.DB {
	Db = connectDB()
	return Db
}

func connectDB() *sql.DB {
	err1 := godotenv.Load()

	if err1 != nil {
		fmt.Printf("Error loading .env file: %v", err1)
		return nil
	}
	DB_USERNAME := os.Getenv("user")
	DB_PASSWORD := os.Getenv("pass")
	DB_NAME := os.Getenv("nume")
	DB_HOST := os.Getenv("host")
	DB_PORT := os.Getenv("port")

	var err error
	ruta := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
	db, err := sql.Open("mysql", ruta)

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)

		return nil
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error verifying connection to database : error=%v", err)

		return nil
	}

	return db
}

func CloseDB(db *sql.DB) {
	defer db.Close()
}
