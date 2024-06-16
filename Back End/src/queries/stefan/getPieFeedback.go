package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetPieFeedback(c *gin.Context) {
	// id_scoala, id_clasa, materie
	var db *sql.DB = database.InitDb()
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din GET
	id_scoala := c.Query("id_scoala")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: id_scoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin in aceasta scoala"})
		return
	}
	id_clasa := c.Query("id_clasa")
	nume_disciplina := c.Query("nume_disciplina")
	q := `SELECT 
    COALESCE(SUM(CASE WHEN tip = 1 THEN 1 ELSE 0 END), 0) AS positive_feedback_count,
    COALESCE(SUM(CASE WHEN tip = 0 THEN 1 ELSE 0 END), 0) AS negative_feedback_count
FROM 
    feedback
WHERE 
    id_scoala = ?
    AND id_clasa = ?
    AND nume_disciplina = ?
    AND directie = 1;`
	var p_count, n_count int
	err := db.QueryRow(q, id_scoala, id_clasa, nume_disciplina).Scan(&p_count, &n_count)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H(map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"type":                  "pie",
				"values":                []interface{}{p_count, n_count},
				"labels":                []interface{}{"Positiv", "Negativ"},
				"textinfo":              "label+percent",
				"insidetextorientation": "radial",
			},
		},
		"layout": map[string]interface{}{},
	}))
}
