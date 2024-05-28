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

type Data struct {
	X    []string  `json:"x"`
	Y    []float64 `json:"y"`
	NUME string    `json:"nume"`
	TIP  string    `json:"type"`
}

func GetBoxPlotMedii(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Extragere date din GET
	idScoala := c.Query("id_scoala")
	//Verificare daca userul este logat
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Verificare daca userul este ADMIN
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}

	// Etape pentru a obține datele necesare
	genuri := map[string][]float64{}
	etnii := map[string][]float64{}
	genList := []string{}
	etnieList := []string{}

	// Obține genurile distincte
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
			genList = append(genList, gen)
			genuri[gen] = []float64{}
		}
	}

	// Obține etniile distincte
	q = `SELECT DISTINCT etnie FROM elev WHERE id_scoala = ?`
	rows, err = db.Query(q, idScoala)
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
			etnieList = append(etnieList, etnie)
			etnii[etnie] = []float64{}
		}
	}

	// Obține mediile generale pentru fiecare elev
	q = `
	SELECT e.id_elev, e.gen, e.etnie, AVG(n.nota) as media_generala
	FROM elev e
	JOIN note n ON e.id_elev = n.id_elev
	WHERE e.id_scoala = ?
	GROUP BY e.id_elev, e.gen, e.etnie`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var idElev int
		var gen string
		var etnie string
		var mediaGenerala float64
		if err := rows.Scan(&idElev, &gen, &etnie, &mediaGenerala); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			genuri[gen] = append(genuri[gen], mediaGenerala)
			etnii[etnie] = append(etnii[etnie], mediaGenerala)
		}
	}

	// Organizează datele pentru plotare
	var dateGenuri []Data = []Data{}
	for _, gen := range genList {
		dateGenuri = append(dateGenuri, Data{
			X:    []string{gen},
			Y:    genuri[gen],
			NUME: gen,
			TIP:  "box",
		})
	}

	var dateEtnii []Data = []Data{}
	for _, etnie := range etnieList {
		dateEtnii = append(dateEtnii, Data{
			X:    []string{etnie},
			Y:    etnii[etnie],
			NUME: etnie,
			TIP:  "box",
		})
	}

	// Returnează datele ca JSON
	c.IndentedJSON(http.StatusOK, gin.H{
		"data_genuri": dateGenuri,
		"data_etnii":  dateEtnii,
		"layout": map[string]interface{}{
			"title": "Repartitia mediilor pe genuri si etnii",
		},
	})
}
