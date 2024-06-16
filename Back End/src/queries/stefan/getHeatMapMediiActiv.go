package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetHeatMapMediiActiv(c *gin.Context) {
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
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var y = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var z = [][]int{}
	for i := 0; i <= 9; i++ {
		z = append(z, []int{})
		for j := 0; j <= 9; j++ {
			z[i] = append(z[i], 0)
		}
	}
	q := `SELECT AVG(medie) AS avg_medie, avg(a.valoare)
FROM (
    SELECT 
        n.id_clasa AS cls, 
        n.id_elev AS elv, 
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
        n.id_clasa, n.id_elev, n.nume_disciplina
) AS subquery  join activitate a on a.id_scoala = ?
and a.id_clasa = cls
and a.id_elev = elv
where not a.valoare = 0 
GROUP BY 
    cls, elv;
`
	rows, err := db.Query(q, idScoala, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var medie, activ float64

		if err := rows.Scan(&medie, &activ); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			i := int(medie) - 1
			j := int(activ) - 1
			z[i][j]++
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"z":           z,
		"x":           x,
		"y":           y,
		"type":        "heatmap",
		"hoverongaps": false,
	}}, "layout": map[string]interface{}{
		"title": "Heatmap medii/activitate",
		"xaxis": map[string]interface{}{
			"title": "Activitate",
		},
		"yaxis": map[string]interface{}{
			"title": "Medie",
		},
	}})
}
