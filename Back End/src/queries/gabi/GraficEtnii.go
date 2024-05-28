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

	for etnie := range etnii {
		q := `
		SELECT e.id_elev, n.materie, n.nota
		FROM elev e
		JOIN note n ON e.id_elev = n.id_elev
		WHERE e.id_scoala = ? AND e.etnie = ?
		`
		rows, err := db.Query(q, idScoala, etnie)
		if err != nil {
			fmt.Println("Eroare: ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
			return
		}
		defer rows.Close()

		eleviNote := make(map[int]map[string][]float64)

		for rows.Next() {
			var idElev int
			var materie string
			var nota float64
			if err := rows.Scan(&idElev, &materie, &nota); err != nil {
				fmt.Println("Eroare: ", err)
			} else {
				if _, exists := eleviNote[idElev]; !exists {
					eleviNote[idElev] = make(map[string][]float64)
				}
				eleviNote[idElev][materie] = append(eleviNote[idElev][materie], nota)
			}
		}

		eleviMedii := make(map[int]float64)

		for idElev, materii := range eleviNote {
			var sumaMedii float64
			for _, note := range materii {
				var sumaNote float64
				for _, nota := range note {
					sumaNote += nota
				}
				mediaMaterie := sumaNote / float64(len(note))
				sumaMedii += mediaMaterie
			}
			mediaGenerala := sumaMedii / float64(len(materii))
			eleviMedii[idElev] = mediaGenerala
		}

		for _, mediaGenerala := range eleviMedii {
			etnii[etnie] = append(etnii[etnie], mediaGenerala)
		}
	}

	var date []DataEtnie
	for etnie, medii := range etnii {
		if len(medii) > 0 {
			date = append(date, DataEtnie{
				X:    medii,
				NAME: etnie,
				TIP:  "box",
			})
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Repartitia mediilor pe etnii",
	}})
}
