package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type session struct {
	nume    string
	prenume string
	id      int
}

var Sessions = map[string]session{}

// Signin adauga un utilizator in DB
func Signup(context *gin.Context) {
	nume := context.PostForm("nume")
	prenume := context.PostForm("prenume")
	email := context.PostForm("email")
	parola, err1 := bcrypt.GenerateFromPassword([]byte(context.PostForm("parola")), 10)

	if err1 != nil {
		fmt.Printf("Eroare: %v", err1)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	}
	var a int
	var q1 string = "insert into cont(nume,prenume,email,parola) values(?,?,?,?)"
	var db *sql.DB = database.InitDb()

	err2 := db.QueryRow(q1, nume, prenume, email, parola).Scan(&a)
	switch {

	case err2 == sql.ErrNoRows:
		context.IndentedJSON(http.StatusOK, "Inserat cu succes!")
	case err2 != nil:
		fmt.Printf("Eroare: %v", err2)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	default:
		//context.IndentedJSON(http.StatusOK, "")
	}
	database.CloseDB(db)
}

func Login(context *gin.Context) {
	email := context.PostForm("email")
	parola := []byte(context.PostForm("parola"))
	var db *sql.DB = database.InitDb()
	var q1 string = "select nume, prenume,id,parola from cont where email=?"
	var hashed []byte
	var nume string
	var prenume string
	var id int
	err1 := db.QueryRow(q1, email).Scan(&nume, &prenume, &id, &hashed)
	switch {
	case err1 == sql.ErrNoRows:
		context.IndentedJSON(http.StatusBadRequest, "No user with this email/password")
		return

	case err1 != nil:
		fmt.Printf("Eroare: %v", err1)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	default:
		if bcrypt.CompareHashAndPassword(hashed, parola) == nil {
			sessionToken := uuid.NewString()
			Sessions[sessionToken] = session{
				nume:    nume,
				prenume: prenume,
				id:      id,
			}
			context.SetCookie("session_cookie", sessionToken, int(time.Now().Add(120*time.Second).Unix()), "/", "", true, true)
			context.Done()
		} else {
			context.IndentedJSON(http.StatusBadRequest, "No user with this email/password")
		}
	}
	database.CloseDB(db)

}

func Logout(context *gin.Context) {

	cookie, err := context.Cookie("session_cookie")
	if err != nil {
		fmt.Println(err)
		return
	}
	delete(Sessions, cookie)
	context.Done()
}
