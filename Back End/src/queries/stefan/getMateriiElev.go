package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetMaterii(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.Query("id_scoala")
	idClasa := c.Query("id_clasa")

	if (!VerificareRol(Rol{
		ROL:    "Elev",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este elev in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este elev in aceasta scoala"})
		return
	}

	q := `select nume_disciplina
		from incadrare
		where id_clasa = ?
		and id_scoala = ?`
	rows, err := db.Query(q, idScoala, idClasa)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var materii = []string{}
	for rows.Next() {
		var materie string
		if err := rows.Scan(&materie); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			materii = append(materii, materie)
		}
	}
	c.IndentedJSON(http.StatusOK, materii)
}
