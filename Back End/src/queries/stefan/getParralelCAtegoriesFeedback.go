package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetParralelCategoriesFeedback(c *gin.Context) {
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
    SUM(CASE WHEN f.tip = 1 THEN 1 ELSE 0 END) AS positive_feedback_count,
    SUM(CASE WHEN f.tip = 0 THEN 1 ELSE 0 END) AS negative_feedback_count
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
    nume ASC;
`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	var prof = []string{}
	var materii = []string{}
	var feedback = []string{}
	var valori = []int{}
	var culori = []int{}
	defer rows.Close()
	for rows.Next() {
		var nume, materie string
		var cnt, poz, neg int

		if err := rows.Scan(&nume, &materie, &cnt, &poz, &neg); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			if poz != 0 {
				prof = append(prof, nume)
				materii = append(materii, materie)
				feedback = append(feedback, "Pozitiv")
				valori = append(valori, poz)
				culori = append(culori, 1)
			}
			if neg != 0 {
				prof = append(prof, nume)
				materii = append(materii, materie)
				feedback = append(feedback, "Negativ")
				valori = append(valori, neg)
				culori = append(culori, 0)
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"type": "parcats",
		"dimensions": []map[string]interface{}{{
			"label":  "Profesor",
			"values": prof,
		}, {
			"label":  "Materie",
			"values": materii,
		}, {
			"label":         "Feedback",
			"values":        feedback,
			"categoryarray": []string{"Negativ", "Pozitiv"},
			"ticktext":      []string{"Negativ", "Pozitiv"},
		}},
		"line": map[string]interface{}{
			"color": culori,
			"colorscale": [][]string{{
				"0", "tomato",
			}, {
				"1", "mediumseagreen",
			}},
		},
		"counts": valori,
	}}, "layout": map[string]interface{}{
		"title": "Feedback profesor materie",
		"width": 600,
	}})
}
