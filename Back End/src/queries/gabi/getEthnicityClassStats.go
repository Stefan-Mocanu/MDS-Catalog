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

type EthnicityStats struct {
	Type  string    `json:"type"`
	R     []float64 `json:"r"`
	Theta []string  `json:"theta"`
	Fill  string    `json:"fill"`
	Class string    `json:"class"`
}

func EthnicityClassStats(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	// Verificăm dacă sesiunea este activă
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extragem id-ul școlii și clasa din query string
	idScoala := c.Query("id_scoala")
	idClasa := c.Query("id_clasa")
	if idScoala == "" || idClasa == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID-ul școlii sau al clasei lipsește"})
		return
	}

	// Verificăm dacă utilizatorul are rolul necesar pentru a accesa aceste date
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator", // Modifică rolul în funcție de cerințe
		ID:     ver,
		SCOALA: idScoala,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Accesul interzis"})
		return
	}

	// Interogare pentru a obține datele necesare pentru graficul de tip scatter polar
	q := `
		SELECT
			etnie,
			AVG(nota) AS avg_nota
		FROM
			note n
		JOIN
			elev e ON n.id_elev = e.id_elev AND n.id_scoala = e.id_scoala AND n.id_clasa = e.id_clasa
		WHERE
			n.id_scoala = ? AND
			n.id_clasa = ?
		GROUP BY
			etnie
		ORDER BY
			etnie
	`
	rows, err := db.Query(q, idScoala, idClasa)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice-uri pentru a stoca datele necesare pentru graficul de tip scatter polar
	var rValues []float64
	var thetaValues []string

	for rows.Next() {
		var etnie string
		var avgNota float64
		if err := rows.Scan(&etnie, &avgNota); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		thetaValues = append(thetaValues, etnie)
		rValues = append(rValues, avgNota)
	}

	// Definim structura de date pentru Plotly
	data := []map[string]interface{}{
		{
			"type":  "scatterpolar",
			"r":     rValues,
			"theta": thetaValues,
			"fill":  "toself",
		},
	}

	// Definim layout-ul pentru graficul de tip scatter polar
	layout := map[string]interface{}{
		"polar": map[string]interface{}{
			"radialaxis": map[string]interface{}{
				"visible": true,
				"range":   [2]float64{0, 10}, // Modifică intervalul în funcție de necesități
			},
		},
		"showlegend": false,
	}

	// Returnăm datele sub formă de răspuns JSON
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
