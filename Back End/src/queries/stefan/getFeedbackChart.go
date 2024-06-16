package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetFeedbackChart(c *gin.Context) {
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
	q := `SELECT 
    CONCAT(p.nume, " ", p.prenume) AS nume,
    i.nume_disciplina AS subject_name,
    COUNT(f.id_feedback) AS feedback_count,
    SUM(CASE WHEN f.tip = 1 and f.directie = 0 THEN 1 ELSE 0 END) AS positive_feedback_count,
    SUM(CASE WHEN f.tip = 0 THEN 1 ELSE 0 END) AS negative_feedback_count,
    SUM(CASE WHEN f.tip = 1 and f.directie = 1 THEN 1 ELSE 0 END) AS positive_feedback_count1,
    SUM(CASE WHEN f.tip = 0 THEN 1 ELSE 1 END) AS negative_feedback_count1
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
    feedback_count DESC, nume ASC;`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var nume = []string{}
	var materii = []string{}
	var poz1 = []int{}
	var poz2 = []int{}
	var neg1 = []int{}
	var neg2 = []int{}
	for rows.Next() {
		var numec, materie string
		var cnt1, p1, p2, n2, n1 int
		if err := rows.Scan(&numec, &materie, &cnt1, &p1, &n1, &p2, &n2); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			nume = append(nume, numec)
			materii = append(materii, materie)
			poz1 = append(poz1, p1)
			poz2 = append(poz2, p2)
			neg1 = append(neg1, n1)
			neg2 = append(neg2, n2)
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"x":           poz1,
		"y":           nume,
		"xaxis":       "x1",
		"yaxis":       "y1",
		"type":        "bar",
		"name":        "Feedback-uri pozitive date de profesor",
		"orientation": "h",
	}, {
		"x":           neg1,
		"y":           nume,
		"xaxis":       "x1",
		"yaxis":       "y1",
		"type":        "bar",
		"name":        "Feedback-uri negative date de profesor",
		"orientation": "h",
	}, {
		"x":     poz2,
		"y":     nume,
		"xaxis": "x2",
		"yaxis": "y1",
		"mode":  "lines+markers",
		"name":  "Feedback-uri pozitive date de elevi",
	}, {
		"x":     neg2,
		"y":     nume,
		"xaxis": "x2",
		"yaxis": "y1",
		"mode":  "lines+markers",
		"name":  "Feedback-uri negative date de elevi",
	},
	}, "layout": map[string]interface{}{
		"margin": map[string]interface{}{
			"r": 25,
			"l": 25,
			"b": 10,
		},
		"width":   600,
		"title":   "Feedback-uri pozitive si negative",
		"barmode": "stack",
	}})
}
