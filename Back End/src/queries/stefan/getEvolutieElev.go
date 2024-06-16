package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetEvolutieElev(c *gin.Context) {
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
	idElev := c.Query("id_elev")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Parinte",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este parinte in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este parinte in aceasta scoala"})
		return
	}
	q := `select 
	nume_disciplina, 
    avg(nota) as media1,
    (SELECT AVG(n1.nota)
     FROM note n1
     WHERE n1.id_scoala = ?
       AND n1.id_clasa = ?
       AND n1.id_elev = ?
       AND n1.nume_disciplina = n.nume_disciplina
    	and n1.data < DATE_SUB(NOW(), INTERVAL 30 DAY)) as media2
	from note n
	where id_scoala = ?
	and id_clasa = ?
	and id_elev=?
	group by 1`
	rows, err := db.Query(q, idScoala, idClasa, idElev, idScoala, idClasa, idElev)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var materii = []string{}
	var medii1 = []float64{}
	var medii2 = []float64{}
	for rows.Next() {
		var materie string
		var m1, m2 float64
		if err := rows.Scan(&materie, &m1, &m2); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			materii = append(materii, materie)
			medii1 = append(medii1, m1)
			medii2 = append(medii2, m2)
		}
	}
	type data struct {
		TYPE   string                 `json:"type"`
		MODE   string                 `json:"mode"`
		VALUE  float64                `json:"value"`
		DELTA  map[string]interface{} `json:"delta"`
		DOMAIN map[string]interface{} `json:"domain"`
		TITLE  map[string]interface{} `json:"title"`
		GAUGE  map[string]interface{} `json:"gauge"`
	}
	date := []data{}
	leng := len(materii)
	step := float64(1) / float64(leng)
	pauza := step / 6
	for i := 0; i < len(materii); i++ {
		date = append(date, data{
			TYPE:  "indicator",
			MODE:  "number+gauge+delta",
			VALUE: medii1[i],
			DELTA: map[string]interface{}{
				"reference": medii2[i],
			},
			DOMAIN: map[string]interface{}{
				"x": []float64{0.25, 1},
				"y": []float64{step*float64(i) + pauza, step*(float64(i)+1) - pauza},
			},
			TITLE: map[string]interface{}{
				"text": materii[i],
			},
			GAUGE: map[string]interface{}{
				"shape": "bullet",
				"axis": map[string]interface{}{
					"range": []int{0, 10},
				},
				"threshold": map[string]interface{}{
					"line": map[string]interface{}{
						"color": "black",
						"width": 2,
					},
					"thickness": 0.75,
					"value":     medii2[i],
				},
				"bar": map[string]interface{}{
					"color": "black",
				},
			},
		})

	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"margin": map[string]interface{}{
			"r": 25,
			"l": 25,
			"b": 10,
		},
		"width": 600,
		"title": "Mediile elevului, cu evolutia din ultima luna",
	}})
}
