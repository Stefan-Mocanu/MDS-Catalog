package gabi

import (
	"backend/database"
	"backend/queries/stefan"
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func HeatMapMediiLunileAnului(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Ensure the database connection is closed

	// Check if the session is active
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extract school ID from query string
	idScoala := c.Query("id_scoala")
	if idScoala == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipsește"})
		return
	}

	// Check if the user has the role of Administrator for the specified school
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este admin în această scoală")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această scoală"})
		return
	}

	// Query to get the average grades per subject and month of the year
	q := `
		SELECT 
			n.nume_disciplina AS materie,
			DATE_FORMAT(n.data, '%Y-%m') AS luna,
			AVG(n.nota) AS avg_medie
		FROM 
			note n
		WHERE 
			n.id_scoala = ?
		GROUP BY 
			n.nume_disciplina, DATE_FORMAT(n.data, '%Y-%m')
		ORDER BY 
			n.nume_disciplina, DATE_FORMAT(n.data, '%Y-%m')
	`

	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Maps to store the subjects, months, and their corresponding average grades
	subjects := make(map[string]struct{})
	months := make(map[string]struct{})
	grades := make(map[string]map[string]float64)

	// Iterating through the results to populate the maps
	for rows.Next() {
		var materie, luna string
		var avgMedie float64
		if err := rows.Scan(&materie, &luna, &avgMedie); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		subjects[materie] = struct{}{}
		months[luna] = struct{}{}
		if _, ok := grades[materie]; !ok {
			grades[materie] = make(map[string]float64)
		}
		grades[materie][luna] = avgMedie
	}

	// Convert maps to slices
	var xValues []string // subjects
	var yValues []string // months
	for subject := range subjects {
		xValues = append(xValues, subject)
	}
	for month := range months {
		yValues = append(yValues, month)
	}

	// Ensure months are sorted in chronological order
	sort.Slice(yValues, func(i, j int) bool {
		ti, _ := time.Parse("2006-01", yValues[i])
		tj, _ := time.Parse("2006-01", yValues[j])
		return ti.Before(tj)
	})

	// Initialize the zValues matrix
	zValues := make([][]float64, len(yValues))
	for i := range zValues {
		zValues[i] = make([]float64, len(xValues))
	}

	// Populate the zValues matrix with grades
	for i, month := range yValues {
		for j, subject := range xValues {
			if avgMedie, ok := grades[subject][month]; ok {
				zValues[i][j] = avgMedie
			} else {
				zValues[i][j] = 0.0 // Default value for missing data
			}
		}
	}

	// Define the color scale for the heatmap
	colorscaleValue := [][]interface{}{
		{0, "#3D9970"},
		{1, "#001f3f"},
	}

	// Build the data for the Plotly heatmap
	data := []map[string]interface{}{
		{
			"x":          xValues,
			"y":          yValues,
			"z":          zValues,
			"type":       "heatmap",
			"colorscale": colorscaleValue,
			"showscale":  false,
		},
	}

	// Define the layout for the heatmap
	layout := map[string]interface{}{
		"title": "Heatmap Medii/Luni Anului",
		"xaxis": map[string]interface{}{
			"title": "Materii",
		},
		"yaxis": map[string]interface{}{
			"title": "Luni",
		},
	}

	// Return the data as a JSON response
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
