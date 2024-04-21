package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Signin(context *gin.Context) {
	nume := context.PostForm("nume")
	prenume := context.PostForm("prenume")
	email := context.PostForm("email")
	parola, err1 := bcrypt.GenerateFromPassword([]byte(context.PostForm("parola")), 10)

	if err1 != nil {
		fmt.Println("Eroare: %v", err1)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	}

	var q1 string = "insert into cont(nume,prenume,email,parola) values(?,?,?,?)"
	var db *sql.DB = database.InitDb()
	var a int
	err2 := db.QueryRow(q1, nume, prenume, email, parola).Scan(&a)
	switch {

	case err2 == sql.ErrNoRows:
		context.IndentedJSON(http.StatusOK, "Inserat cu succes!")
	case err2 != nil:
		fmt.Println("Eroare: %v", err2)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	default:
		//context.IndentedJSON(http.StatusOK, a)
	}
	database.CloseDB(db)
}

func Login(context *gin.Context) {

}

func Logout(context *gin.Context) {

}
