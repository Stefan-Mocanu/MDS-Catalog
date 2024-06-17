package gabi

import (
	"backend/database"
	"backend/queries/stefan"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type ClasaStats struct {
	X    []int  `json:"x"`
	Y    []int  `json:"y"`
	MODE string `json:"mode"`
	TIP  string `json:"type"`
}

func EleviPromov(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Ensure database connection closure

	// Check if session is active
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extract school ID from query string
	idScoalaStr := c.Query("id_scoala")
	if idScoalaStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipseste"})
		return
	}
	idScoala, err := strconv.Atoi(idScoalaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala invalid"})
		return
	}

	// Check if the user has the Administrator role for the specified school
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoalaStr,
		ID:     ver,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această școală"})
		return
	}

	// Query to get the percentage of students who passed (average grade over 5) for each class
	q := `
		SELECT e.id_clasa, COUNT(*) * 100.0 / (SELECT COUNT(*) FROM elev WHERE id_clasa = e.id_clasa AND id_scoala = ?) AS procent_trecere
		FROM note n
		JOIN elev e ON n.id_elev = e.id_elev AND n.id_scoala = e.id_scoala AND n.id_clasa = e.id_clasa
		WHERE n.id_scoala = ? AND n.nota > 5
		GROUP BY e.id_clasa
		ORDER BY e.id_clasa`
	log.Printf("Executing query: %s with idScoala = %d", q, idScoala)
	rows, err := db.Query(q, idScoala, idScoala)
	if err != nil {
		log.Printf("Database query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slices to store data for the dot plot
	var xValues []string
	var yValues []int

	for rows.Next() {
		var idClasa string
		var procent float64
		if err := rows.Scan(&idClasa, &procent); err != nil {
			log.Printf("Error scanning results: %v", err)
			continue
		}
		xValues = append(xValues, idClasa)
		yValues = append(yValues, int(procent))
	}

	// Check if any data was retrieved
	if len(xValues) == 0 {
		log.Println("No data found for the specified school.")
		c.JSON(http.StatusOK, gin.H{"data": nil, "message": "Nu s-au găsit date pentru școala specificată."})
		return
	}

	// Data for the dot plot in JSON format
	data := []map[string]interface{}{
		{
			"x":    xValues,
			"y":    yValues,
			"mode": "markers",
			"type": "scatter",
			"name": "Percent of students passing with average > 5",
			"marker": map[string]interface{}{
				"color": "rgba(156, 165, 196, 0.95)",
				"line": map[string]interface{}{
					"color": "rgba(156, 165, 196, 1.0)",
					"width": 1,
				},
				"symbol": "circle",
				"size":   16,
			},
		},
	}

	log.Printf("Query successful, returning data: %v", data)
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": gin.H{
		"barmode": "stack",
	}})
}
