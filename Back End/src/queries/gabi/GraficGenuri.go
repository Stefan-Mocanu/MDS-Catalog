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

type DataGen struct {
	X    []float64 `json:"x"`
	NAME string    `json:"name"`
	TIP  string    `json:"type"`
}

func GetDistGenuri(c *gin.Context) {
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

	genuri := map[string][]float64{}

	q := `SELECT DISTINCT gen FROM elev WHERE id_scoala = ?`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var gen string
		if err := rows.Scan(&gen); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			genuri[gen] = []float64{}
		}
	}

	var date []DataGen
	for gen := range genuri {
		q := `
		SELECT AVG(n.nota) as media_generala
		FROM elev e
		JOIN note n
		where e.id_elev = n.id_elev
		and e.id_clasa = n.id_clasa
		and e.id_scoala = ? AND e.etnie = ?
		GROUP BY e.id_clasa,e.id_elev`
		rows, err := db.Query(q, idScoala, gen)
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
				genuri[gen] = append(genuri[gen], mediaGenerala)
			}
		}
		if len(genuri[gen]) > 0 {
			date = append(date, DataGen{
				X:    genuri[gen],
				NAME: gen,
				TIP:  "box",
			})
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Repartitia mediilor pe genuri",
	}})
}
