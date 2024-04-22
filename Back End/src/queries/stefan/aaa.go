package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func STEVE(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	var a int
	err := db.QueryRow("select count(*) a from cont").Scan(&a)
	switch {

	case err == sql.ErrNoRows:
		fmt.Printf("Eroare: %v", err)

		context.IndentedJSON(http.StatusInternalServerError, 500)
	case err != nil:
		fmt.Printf("Eroare: %v", err)
		context.IndentedJSON(http.StatusInternalServerError, 500)
	default:
		context.IndentedJSON(http.StatusOK, a)
	}
	database.CloseDB(db)
}
