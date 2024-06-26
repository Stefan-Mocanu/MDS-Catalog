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

type Rol struct {
	ROL    string `json:"rol"`
	SCOALA string `json:"scoala"`
	ID     int    `json:"id"`
}

func GetRoluri(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	var roluri []Rol = []Rol{}
	cookie, err := context.Cookie("session_cookie")
	if err != nil {
		context.IndentedJSON(http.StatusOK, false)
		return
	}
	content, ok := Sessions[cookie]
	if !ok {
		context.IndentedJSON(http.StatusOK, false)
		return
	}
	q := "select id_rol, s.nume nume, c.id_scoala scoala from cont_rol c, scoala s where s.id = c.id_scoala and id_cont=?"
	rows, err1 := db.Query(q, content.ID)
	if err1 != nil {
		fmt.Println("Eroare: ", err1)
		context.IndentedJSON(http.StatusOK, false)
		return
	}
	for rows.Next() {
		var aux Rol
		if err := rows.Scan(&aux.ROL, &aux.SCOALA, &aux.ID); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			roluri = append(roluri, aux)
		}
	}
	defer rows.Close()
	context.IndentedJSON(http.StatusOK, roluri)
	database.CloseDB(db)
}
