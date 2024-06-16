package gabi

import (
	"backend/database"
	"backend/queries/stefan"
	"database/sql"
	"fmt"
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

	// Extragem id-ul școlii din query string
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

	// Verificăm dacă utilizatorul are rolul de Administrator pentru școala specificată
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoalaStr,
		ID:     ver,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această școală"})
		return
	}

	// Query pentru a obține procentul de elevi care au trecut (media peste 5) pentru fiecare clasă
	q := `
		SELECT id_clasa, COUNT(*) * 100.0 / (SELECT COUNT(*) FROM elev WHERE id_clasa = c.id_clasa AND id_scoala = ?) AS procent_trecere
		FROM note n
		JOIN elev e ON n.id_elev = e.id_elev AND n.id_scoala = e.id_scoala AND n.id_clasa = e.id_clasa
		WHERE n.id_scoala = ? AND n.nota > 5
		GROUP BY id_clasa
		ORDER BY id_clasa
	`
	rows, err := db.Query(q, idScoala, idScoala)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice-uri pentru a stoca datele necesare pentru dot plot
	var xValues []int
	var yValues []int

	for rows.Next() {
		var idClasa int
		var procent float64
		if err := rows.Scan(&idClasa, &procent); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		xValues = append(xValues, idClasa)
		yValues = append(yValues, int(procent))
	}

	// Datele pentru dot plot în format JSON
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

	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": gin.H{
		"barmode": "stack",
	}})
}
