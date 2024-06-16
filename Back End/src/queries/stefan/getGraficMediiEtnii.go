package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetGraficMediiEtnii(c *gin.Context) {
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
	etnii := map[string][]float64{}
	absente_etnii := map[string][]int{}
	q := `SELECT DISTINCT etnie FROM elev WHERE id_scoala = ?`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var etnie string
		if err := rows.Scan(&etnie); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			etnii[etnie] = []float64{}
			absente_etnii[etnie] = []int{}
		}
	}
	q = `SELECT etn, AVG(medie) AS avg_medie
FROM (
    SELECT 
        n.id_clasa AS cls, 
        n.id_elev AS elv, 
        e.etnie AS etn,
        AVG(n.nota) AS medie
    FROM 
        note n
    JOIN 
        elev e ON n.id_scoala = e.id_scoala 
              AND n.id_clasa = e.id_clasa 
              AND n.id_elev = e.id_elev
    WHERE 
        n.id_scoala = ?
    GROUP BY 
        n.id_clasa, n.id_elev, e.etnie, n.nume_disciplina
) AS subquery
GROUP BY 
    cls, elv, etn;`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nota float64
		var etnie string
		if err := rows.Scan(&etnie, &nota); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			etnii[etnie] = append(etnii[etnie], nota)
		}
	}
	q = `select e.etnie, count(*)
from activitate a 
join elev e
on e.id_scoala = ?
and e.id_scoala = a.id_scoala
and e.id_clasa = a.id_clasa
and e.id_elev = a.id_elev
where a.valoare = 0 
group by a.id_clasa, a.id_elev, e.etnie`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nr int
		var etnie string
		if err := rows.Scan(&etnie, &nr); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			absente_etnii[etnie] = append(absente_etnii[etnie], nr)
		}
	}
	type data struct {
		X    []float64 `json:"x"`
		Y    []int     `json:"y"`
		MODE string    `json:"mode"`
		NAME string    `json:"name"`
		TYPE string    `json:"type"`
	}
	var date = []data{}
	for etnie := range etnii {
		date = append(date, data{
			X:    etnii[etnie],
			Y:    absente_etnii[etnie],
			MODE: "markers",
			NAME: etnie,
			TYPE: "scatter",
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"title": "Medii si absente pe etnii",
		"xaxis": map[string]interface{}{
			"title": "Medii",
		},
		"yaxis": map[string]interface{}{
			"title": "Nr. absente",
		},
	}})
}
