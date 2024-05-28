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
	date := []DataEtnie{}
	for etnie := range etnii {
		q := `
		select round(avg(a.mean),2) mean
		from (SELECT n.id_clasa,n.id_elev,avg(n.nota) mean
				from note n, elev e
				where e.id_scoala = ?
			and e.id_scoala = n.id_scoala
			and e.id_clasa = n.id_clasa
			and e.id_elev = n.id_elev
			and e.etnie = ?
				GROUP by nume_disciplina, id_clasa, id_elev) a
		GROUP by a.id_clasa, a.id_elev
		`
		rows, err := db.Query(q, idScoala, etnie)
		if err != nil {
			fmt.Println("Eroare: ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var nota float64

			if err := rows.Scan(&nota); err != nil {
				fmt.Println("Eroare: ", err)
			} else {
				etnii[etnie] = append(etnii[etnie], nota)
			}

		}
		date = append(date, DataEtnie{
			X:    etnii[etnie],
			NAME: etnie,
			TIP:  "box",
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Repartitia mediilor pe etnii",
	}})
}

//
