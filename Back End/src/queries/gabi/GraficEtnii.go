package gabi

import (
	"backend/database"
	"backend/queries/stefan"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type DataEtnie struct {
	X     string  `json:"x"`
	MEDIE float64 `json:"medie"`
	TIP   string  `json:"type"`
}

func GetDistEtnii(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	idScoala := c.Query("id_scoala")
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}

	etnii := map[string][]int{}

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
			etnii[etnie] = []int{}
		}
	}

	var date []DataEtnie
	for etnie := range etnii {
		q := `
		SELECT AVG(n.nota) as media_generala
		FROM elev e
		JOIN note n ON e.id_elev = n.id_elev
		WHERE e.id_scoala = ? AND e.etnie = ?
		GROUP BY e.id_elev`
		rows, err := db.Query(q, idScoala, etnie)
		if err != nil {
			fmt.Println("Eroare: ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
			return
		}
		defer rows.Close()

		var sumaMedii float64
		var count int
		for rows.Next() {
			var mediaGenerala float64
			if err := rows.Scan(&mediaGenerala); err != nil {
				fmt.Println("Eroare: ", err)
			} else {
				sumaMedii += mediaGenerala
				count++
			}
		}
		if count > 0 {
			mediaFinala := sumaMedii / float64(count)
			date = append(date, DataEtnie{
				X:     etnie,
				MEDIE: mediaFinala,
				TIP:   "bar",
			})
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"barmode": "stack",
	}})
}
