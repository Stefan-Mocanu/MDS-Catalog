package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type elev struct {
	CLASA   string `json:"clasa"`
	ID      int    `json:"id"`
	NUME    string `json:"nume"`
	PRENUME string `json:"prenume"`
}

func GetElevi(c *gin.Context) {
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
	//Verificare daca useul este PARINTE
	if (!VerificareRol(Rol{
		ROL:    "Parinte",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este parinte in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este parinte in aceasta scoala"})
		return
	}
	//Obtinere elevi
	q := `select id_clasa, id_elev, nume, prenume
		from elev
		where id_cont_parinte = ?`
	rows, err := db.Query(q, ver)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var elevi = []elev{}
	for rows.Next() {
		var clasa, nume, prenume string
		var id int
		if err := rows.Scan(&clasa, &id, &nume, &prenume); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			fmt.Println(id)
			elevi = append(elevi, elev{
				CLASA:   clasa,
				ID:      id,
				NUME:    nume,
				PRENUME: prenume,
			})
		}
	}
	c.IndentedJSON(http.StatusOK, elevi)

	database.CloseDB(db)
}
