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

	date := []DataEtnie{}
	for gen := range genuri {
		q := `
		select round(avg(a.mean),2) mean
		from (SELECT n.id_clasa,n.id_elev,avg(n.nota) mean
				from note n, elev e
				where e.id_scoala = ?
			and e.id_scoala = n.id_scoala
			and e.id_clasa = n.id_clasa
			and e.id_elev = n.id_elev
			and e.gen = ?
				GROUP by nume_disciplina, id_clasa, id_elev) a
		GROUP by a.id_clasa, a.id_elev
		`
		rows, err := db.Query(q, idScoala, gen)
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
				genuri[gen] = append(genuri[gen], nota)
			}

		}
		date = append(date, DataEtnie{
			X:    genuri[gen],
			NAME: gen,
			TIP:  "box",
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Repartitia mediilor pe genuri",
	}})
}
