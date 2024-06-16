package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCorelatieActivNote(c *gin.Context) {
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
	nume_disciplina := c.Query("materie")
	start1 := c.Query("start_date_activ")
	start2 := c.Query("start_date_note")
	end1 := c.Query("end_date_activ")
	end2 := c.Query("end_date_note")
	q := `SELECT id_elev, valoare, data
	FROM activitate
	WHERE id_scoala = ?
	AND NOT valoare=0
	AND id_clasa = ?
	AND nume_disciplina = ?
	AND data BETWEEN ? AND ?;`
	rows, err := db.Query(q, id_scoala, id_clasa, nume_disciplina, start1, end1)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	a_elev := []int{}
	a_val := []int{}
	a_data := []string{}
	for rows.Next() {
		var elev, val int
		var data string
		if err := rows.Scan(&elev, &val, &data); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			a_elev = append(a_elev, elev)
			a_val = append(a_val, val)
			a_data = append(a_data, data)
		}
	}
	q2 := `SELECT id_elev, nota, data
	FROM note
	WHERE id_scoala = ?
	AND id_clasa = ?
	AND nume_disciplina = ?
	AND data BETWEEN ? AND ?;`
	rows2, err := db.Query(q2, id_scoala, id_clasa, nume_disciplina, start2, end2)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows2.Close()
	n_elev := []int{}
	n_val := []int{}
	n_data := []string{}
	for rows2.Next() {
		var elev, val int
		var data string
		if err := rows2.Scan(&elev, &val, &data); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			n_elev = append(n_elev, elev)
			n_val = append(n_val, val)
			n_data = append(n_data, data)
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": []interface{}{
			map[string]interface{}{
				"x":    a_data,
				"y":    a_val,
				"mode": "markers+text",
				"type": "scatter",
				"marker": map[string]interface{}{
					"size": 20,
				},
				"name": "Activitate",
				"text": a_elev,
			},
			map[string]interface{}{
				"x":    n_data,
				"y":    n_val,
				"mode": "markers+text",
				"type": "scatter",
				"marker": map[string]interface{}{
					"size": 20,
				},
				"name": "Note",
				"text": n_elev,
			},
		},
		"layout": map[string]interface{}{
			"legend": map[string]interface{}{
				"yref": "paper",
				"font": map[string]interface{}{
					"family": "Arial, sans-serif",
					"size":   20,
					"color":  "grey",
				},
			},
			"title": "Activitate-Nota",
		},
	})
}
