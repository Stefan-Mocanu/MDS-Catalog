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

type MedieClasa struct {
	Clasa string  `json:"clasa"`
	Medie float64 `json:"medie"`
}

func MediiClase(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Ensure the database connection is closed

	// Check if the session is active
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Obținerea id_scoala, id_clasa și materie din context
	idScoala := c.Query("id_scoala")
	if idScoala == "" {
		fmt.Println("ID scoala lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipseste"})
		return
	}

	idClasa := c.Query("id_clasa")
	if idClasa == "" {
		fmt.Println("ID clasa lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID clasa lipseste"})
		return
	}

	materie := c.Query("materie")
	if materie == "" {
		fmt.Println("Materie lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Materie lipseste"})
		return
	}

	// Verify if the user has the role of Professor
	if !stefan.VerificareRol(stefan.Rol{
		ROL: "Profesor",
		SCOALA: idScoala,
		ID:  ver,
	}) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor"})
		return
	}

	// Query to get the average grade for the specified class and subject
	q := `
		SELECT n.id_clasa, AVG(n.nota) AS medie
		FROM note n
		JOIN incadrare i ON n.id_scoala = i.id_scoala AND n.id_clasa = i.id_clasa AND n.nume_disciplina = i.nume_disciplina
		WHERE n.id_scoala = ? AND n.id_clasa = ? AND n.nume_disciplina = ?
		GROUP BY n.id_clasa
	`
	rows, err := db.Query(q, idScoala, idClasa, materie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice to store the class averages
	var mediiClase []MedieClasa

	for rows.Next() {
		var medieClasa MedieClasa
		if err := rows.Scan(&medieClasa.Clasa, &medieClasa.Medie); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		mediiClase = append(mediiClase, medieClasa)
	}

	// Data for the chart
	var xValues []string
	var yValues []float64

	for _, medieClasa := range mediiClase {
		xValues = append(xValues, medieClasa.Clasa)
		yValues = append(yValues, medieClasa.Medie)
	}

	// Return data in a JSON format compatible with Plotly
	data := []map[string]interface{}{
		{
			"type":     "funnel",
			"name":     "Medii Clase",
			"y":        xValues, // Use xValues as y to display the averages on the axis
			"x":        yValues, // Averages will be the values for the x-axis in the chart
			"textinfo": "value+percent initial",
		},
	}

	layout := gin.H{
		"margin":     gin.H{"l": 130, "r": 0},
		"width":      600,
		"funnelmode": "stack",
		"showlegend": true,
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
