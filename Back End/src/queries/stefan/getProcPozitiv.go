package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetProcPozitiv(c *gin.Context) {
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
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin in aceasta scoala"})
		return
	}
	q := `select t.nume, t.feedback_count, sum(t.pos1)
from(
SELECT 
    CONCAT(p.nume, " ", p.prenume) AS nume,
    i.nume_disciplina AS subject_name,
    COUNT(f.id_feedback) AS feedback_count,
    SUM(CASE WHEN f.tip = 1 and f.directie = 0 THEN 1 ELSE 0 END) AS pos1
FROM 
    profesor p
JOIN 
    incadrare i ON p.id_scoala = i.id_scoala AND p.id = i.id_profesor
LEFT JOIN 
    feedback f ON i.id_scoala = f.id_scoala AND i.id_clasa = f.id_clasa AND i.nume_disciplina = f.nume_disciplina
WHERE 
    p.id_scoala = ?
GROUP BY 
    p.nume, p.prenume, i.nume_disciplina
ORDER BY 
    feedback_count DESC, nume ASC) as t
group by t.nume;`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var nume = []string{}
	var y = []float64{}
	for rows.Next() {
		var numec string
		var p1, c1 float64
		if err := rows.Scan(&numec, &c1, &p1); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			nume = append(nume, numec)
			y = append(y, p1/c1)
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"x":     nume,
		"y":     y,
		"xaxis": "x1",
		"yaxis": "y1",
		"type":  "bar",
		"name":  "Procent feedback-uri pozitive primite de profesor",
	},
	}, "layout": map[string]interface{}{
		"title": "Procent feedback-uri pozitive primite de profesor",
	}})
}
