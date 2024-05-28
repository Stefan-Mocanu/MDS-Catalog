package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gonum.org/v1/gonum/stat"
)

type Data2 struct {
	X   []string  `json:"x"`
	Y   []float64 `json:"y"`
	TIP string    `json:"type"`
}

func GetDevStdNote(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Extragere date din GET
	idScoala := c.Query("id_scoala")
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Verifiare daca userul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}
	medii := map[string][]float64{}
	clase := []string{}
	q := `select nume
		from clasa
		where id_scoala = ?`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var clasa string
		if err := rows.Scan(&clasa); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			medii[clasa] = []float64{}
			clase = append(clase, clasa)
		}
	}
	if len(clase) == 0 {
		c.IndentedJSON(http.StatusOK, false)
		return
	}
	q = `select a.id_clasa, round(avg(a.mean),2) mean
			from (SELECT id_clasa,id_elev,avg(nota) mean
					from note
					where id_scoala = ?
					GROUP by nume_disciplina, id_clasa, id_elev) a
			GROUP by a.id_clasa, a.id_elev`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var clasa string
		var medie float64
		if err := rows.Scan(&clasa, &medie); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			medii[clasa] = append(medii[clasa], medie)
		}
	}
	var stDev []float64 = []float64{}
	for key := range clase {
		elem := stat.StdDev(medii[clase[key]], nil)
		if math.IsNaN(elem) {
			elem = 0
		}
		stDev = append(stDev, elem)
	}
	var date []Data2 = []Data2{}
	fmt.Println(stDev)
	fmt.Println(clase)
	date = append(date, Data2{
		X:   clase,
		Y:   stDev,
		TIP: "bar",
	})

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Deviatia standard pentru clase",
	}})
}
