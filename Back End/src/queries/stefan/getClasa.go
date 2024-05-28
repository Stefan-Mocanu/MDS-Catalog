package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetClasa(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din GET
	idScoala := c.Query("id_scoala")
	//Verificare daca userul este ELEV
	if (!VerificareRol(Rol{
		ROL:    "Elev",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este elev in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este elev in aceasta scoala"})
		return
	}
	//Obtinere clase
	q := `select id_clasa
		from elev
		where id_cont_elev = ?`
	rows, err := db.Query(q, ver)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var clase []string = []string{}
	for rows.Next() {
		var clasa string
		if err := rows.Scan(&clasa); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			clase = append(clase, clasa)
		}
	}
	c.IndentedJSON(http.StatusOK, clase)

	database.CloseDB(db)
}
