package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetHeatMapIncadrare(c *gin.Context) {
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
	var x = []string{}
	var y = []string{}
	var z = [][]int{}
	q := `select nume
		from clasa
		where id_scoala=?
		order by nume`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var nume string

		if err := rows.Scan(&nume); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			x = append(x, nume)
		}
	}
	q = `select nume from discipline where id_scoala=? order by nume`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var nume string

		if err := rows.Scan(&nume); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			y = append(y, nume)
		}
	}
	for i := 0; i < len(x); i++ {
		z = append(z, []int{})
		for j := 0; j < len(y); j++ {
			q = `select count(*)
				from incadrare i
				where i.id_scoala = ?
				and i.id_clasa = ?
				and i.nume_disciplina = ?`
			var rez int
			err = db.QueryRow(q, idScoala, x[i], y[j]).Scan(&rez)
			if err != nil {
				fmt.Println("Eroare: ", err)
			} else {
				z[i] = append(z[i], rez)
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"z":           z,
		"x":           x,
		"y":           y,
		"type":        "heatmap",
		"hoverongaps": false,
	}, "layout": map[string]interface{}{
		"title": "Verificare clasa are materie",
		"xaxis": map[string]interface{}{
			"title": "Clase",
		},
		"yaxis": map[string]interface{}{
			"title": "Materie",
		},
	}})
}
