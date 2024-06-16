package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetFeedbackPoints(c *gin.Context) {
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
	idClasa := c.Query("id_clasa")
	materie := c.Query("materie")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}
	q := `SELECT 
    AVG(n.nota) AS average_grade,
    COUNT(f.id_feedback) AS total_feedbacks,
    SUM(CASE WHEN f.tip = 1 THEN 1 ELSE 0 END) - SUM(CASE WHEN f.tip = 0 THEN 1 ELSE 0 END) AS feedback_balance
FROM 
    elev e
LEFT JOIN 
    note n ON e.id_scoala = n.id_scoala AND e.id_clasa = n.id_clasa AND e.id_elev = n.id_elev
LEFT JOIN 
    feedback f ON e.id_scoala = f.id_scoala AND e.id_clasa = f.id_clasa AND e.id_elev = f.id_elev AND f.nume_disciplina = n.nume_disciplina
WHERE 
    e.id_scoala = ?
    AND e.id_clasa = ?
    AND n.nume_disciplina = ?
GROUP BY 
    e.id_elev, e.nume, e.prenume;
`
	rows, err := db.Query(q, idScoala, idClasa, materie)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var medii = []float64{}
	var feedback = []int{}
	var feedback1 = []int{}
	var text = []string{}
	for rows.Next() {
		var medie float64
		var x, y int
		if err := rows.Scan(&medie, &x, &y); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			medii = append(medii, medie)
			feedback = append(feedback, x)
			feedback1 = append(feedback1, y)
			text = append(text, "Medie: "+fmt.Sprint(medie))
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"x":    feedback,
		"y":    feedback1,
		"mode": "markers",
		"marker": map[string]interface{}{
			"size":  40,
			"color": medii,
		},
		"text":      text,
		"hoverinfo": "text",
	}}, "layout": map[string]interface{}{
		"title": "Mediile elevilor care au dat feedback",
		"xaxis": map[string]interface{}{
			"title": "Numar de feedback-uri date",
		},
		"yaxis": map[string]interface{}{
			"title": "Feedback-uri pozitive - Feedback-uri negative",
		},
	}})
}
