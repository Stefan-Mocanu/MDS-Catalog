package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetMedieClasament(c *gin.Context) {
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
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}
	q := `WITH Student1AvgGrades AS (
    SELECT 
        nume_disciplina, 
        AVG(nota) AS avg_nota
    FROM note
    WHERE id_scoala = ?
        AND id_clasa = ?
        AND id_elev = ?
    GROUP BY nume_disciplina
),
AllStudentAvgGrades AS (
    SELECT 
        nume_disciplina,
        id_elev,
        AVG(nota) AS avg_nota
    FROM note
    WHERE id_scoala = ?
        AND id_clasa = ?
    GROUP BY nume_disciplina, id_elev
),
StudentsAboveAvg AS (
    SELECT 
        a.nume_disciplina,
        COUNT(*) AS num_students_above_avg
    FROM AllStudentAvgGrades a
    JOIN Student1AvgGrades s1
        ON a.nume_disciplina = s1.nume_disciplina
    WHERE a.avg_nota > s1.avg_nota
    GROUP BY a.nume_disciplina
)

SELECT 
    s1.nume_disciplina, 
    s1.avg_nota AS student1_avg_nota, 
    COALESCE(sa.num_students_above_avg, 0) +1 AS num_students_above_avg
FROM Student1AvgGrades s1
LEFT JOIN StudentsAboveAvg sa
    ON s1.nume_disciplina = sa.nume_disciplina;
`
	rows, err := db.Query(q, idScoala, idClasa, idElev, idScoala, idClasa)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	var discipline = []string{}
	var medii = []float64{}
	var clasament = []int{}
	for rows.Next() {
		var disc string
		var medie float64
		var clas int
		if err := rows.Scan(&disc, &medie, &clas); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			discipline = append(discipline, disc)
			medii = append(medii, medie)
			clasament = append(clasament, clas)
		}
	}
	type a struct {
		VISIBLE bool  `json:"visible"`
		RANGE   []int `json:"range"`
	}
	type axis struct {
		AXIS a `json:"axis"`
	}
	type b struct {
		COLUMN int `json:"column"`
		ROW    int `json:"row"`
	}
	type gauge struct {
		TYPE   string  `json:"type"`
		VALUE  float64 `json:"value"`
		GAUGE  axis    `json:"gauge"`
		MODE   string  `json:"mode"`
		TITLE  string  `json:"title"`
		DOMAIN b       `json:"domain"`
	}
	var data = []gauge{}
	for i := range discipline {
		data = append(data, gauge{
			TYPE:  "indicator",
			VALUE: medii[i],
			GAUGE: axis{
				AXIS: a{
					VISIBLE: false,
					RANGE:   []int{0, 10},
				},
			},
			TITLE: "Media " + discipline[i],
			DOMAIN: b{COLUMN: 0,
				ROW: i,
			},
		})
		data = append(data, gauge{
			TYPE:  "indicator",
			MODE:  "number",
			TITLE: "Loc in clasament",
			VALUE: float64(clasament[i]),
			DOMAIN: b{COLUMN: 1,
				ROW: i,
			},
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": map[string]interface{}{
		"width":  600,
		"margin": map[string]interface{}{"t": 25, "b": 25, "l": 25, "r": 25},
		"grid":   map[string]interface{}{"rows": len(medii), "columns": 2, "pattern": "independent"},
		"template": map[string]interface{}{
			"data": map[string]interface{}{
				"indicator": []map[string]interface{}{
					{"mode": "number+gauge"},
				},
			},
		},
	}})
}
