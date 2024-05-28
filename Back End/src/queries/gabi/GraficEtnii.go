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
	X    []float64 `json:"x"`
	NAME string    `json:"name"`
	TIP  string    `json:"type"`
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

	etnii := map[string][]float64{}

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
		}
	}

	var date []DataEtnie
	for etnie := range etnii {
		q := `
		SELECT avg(n.nota) as media_generala
		FROM elev e
		JOIN note n
		where e.id_elev = n.id_elev
		and e.id_clasa = n.id_clasa
		and e.id_scoala = ? AND e.etnie = ?
		group by e.id_clasa, e.id_elev
		`
		rows, err := db.Query(q, idScoala, etnie)
		if err != nil {
			fmt.Println("Eroare: ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var mediaGenerala float64
			if err := rows.Scan(&mediaGenerala); err != nil {
				fmt.Println("Eroare: ", err)
			} else {
				etnii[etnie] = append(etnii[etnie], mediaGenerala)
			}
		}
		if len(etnii[etnie]) > 0 {
			date = append(date, DataEtnie{
				X:    etnii[etnie],
				NAME: etnie,
				TIP:  "box",
			})
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Repartitia mediilor pe etnii",
	}})
}
