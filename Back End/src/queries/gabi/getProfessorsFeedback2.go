package gabi

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Structura pentru datele de feedback ale profesorilor
type ProfessorFeedbackStats2 struct {
	X       []float64 `json:"x"`
	Y       []float64 `json:"y"`
	Text    []string  `json:"text"`
	Mode    string    `json:"mode"`
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Marker  Marker2   `json:"marker"`
	ShowLeg bool      `json:"showlegend"`
}

// Structura pentru marker
type Marker2 struct {
	Size int `json:"size"`
}

func ProfessorsFeedback2(c *gin.Context) {
	// Inițializarea conexiunii la baza de date
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	ver := 1
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	idScoalaStr := "1"
	if idScoalaStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipsește"})
		return
	}
	idScoala, err := strconv.Atoi(idScoalaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala invalid"})
		return
	}

	if false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această școală"})
		return
	}

	q := `
		SELECT 
			p.id,
			p.nume,
			IFNULL(SUM(CASE WHEN f.tip THEN 1 ELSE 0 END) * 100.0 / COUNT(f.id_feedback), 0) AS procent_feedback_pozitiv,
			AVG(n.nota) AS media_clase
		FROM 
			profesor p
		LEFT JOIN 
			incadrare i ON p.id = i.id_profesor AND p.id_scoala = i.id_scoala
		LEFT JOIN 
			clasa c ON i.id_clasa = c.nume AND i.id_scoala = c.id_scoala
		LEFT JOIN 
			feedback f ON i.id_scoala = f.id_scoala
		LEFT JOIN 
			note n ON i.id_scoala = n.id_scoala AND c.nume = n.id_clasa
		WHERE 
			p.id_scoala = ?
		GROUP BY 
			p.id
		ORDER BY 
			p.nume
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	var xValues []float64
	var yValues []float64
	var textValues []string

	for rows.Next() {
		var idProfesor int
		var numeProfesor string
		var procentFeedbackPozitiv, mediaClase sql.NullFloat64

		if err := rows.Scan(&idProfesor, &numeProfesor, &procentFeedbackPozitiv, &mediaClase); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}

		if procentFeedbackPozitiv.Valid && mediaClase.Valid {
			xValues = append(xValues, procentFeedbackPozitiv.Float64)
			yValues = append(yValues, mediaClase.Float64)
			textValues = append(textValues, fmt.Sprintf("%s: %.2f%%", numeProfesor, procentFeedbackPozitiv.Float64))
		}
	}

	// Definirea primului trace (Team A)
	trace1 := ProfessorFeedbackStats2{
		X:       xValues,
		Y:       yValues,
		Text:    textValues,
		Mode:    "markers+text",
		Type:    "scatter",
		Marker:  Marker2{Size: 12},
		ShowLeg: true,
	}

	// Definirea al doilea trace (Team B) - exemplu pentru diversitate
	trace2 := ProfessorFeedbackStats2{
		X:       []float64{1.5, 2.5, 3.5, 4.5, 5.5},
		Y:       []float64{4, 1, 7, 1, 4},
		Text:    []string{"B-a", "B-b", "B-c", "B-d", "B-e"},
		Mode:    "markers+text",
		Type:    "scatter",
		Marker:  Marker2{Size: 12},
		ShowLeg: true,
	}

	data := []ProfessorFeedbackStats2{trace1, trace2}

	layout := gin.H{
		"xaxis": gin.H{
			"range": [2]interface{}{0.75, 5.25},
		},
		"yaxis": gin.H{
			"range": [2]interface{}{0, 8},
		},
		"legend": gin.H{
			"y":    0.5,
			"yref": "paper",
			"font": gin.H{
				"family": "Arial, sans-serif",
				"size":   20,
				"color":  "grey",
			},
		},
		"title": "Data Labels on the Plot",
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
