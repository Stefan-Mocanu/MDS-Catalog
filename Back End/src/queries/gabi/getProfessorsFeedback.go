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

type ProfessorFeedbackStats struct {
	X       []float64  `json:"x"`
	Y       []float64  `json:"y"`
	Text    [][]string `json:"text"`
	Mode    string     `json:"mode"`
	Type    string     `json:"type"`
	Name    string     `json:"name"`
	Marker  Marker     `json:"marker"`
	ShowLeg bool       `json:"showlegend"`
}

type Marker struct {
	Size int `json:"size"`
}

func ProfessorsFeedback(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipsește"})
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

	// Interogare pentru a obține procentul de feedback pozitiv și media claselor pentru fiecare profesor
	q := `
	SELECT 
		p.id AS id_profesor,
		p.nume AS nume_profesor,
		IFNULL(SUM(IF(f.tip = 1, 1, 0)) * 100.0 / COUNT(f.id_feedback), 0) AS procent_feedback_pozitiv,
		IFNULL(AVG(n.nota), 0) AS media_clase
	FROM 
		profesor p
	LEFT JOIN 
		incadrare i ON p.id = i.id_profesor AND p.id_scoala = i.id_scoala
	LEFT JOIN 
		feedback f ON i.id_scoala = f.id_scoala AND i.nume_disciplina = f.nume_disciplina
	LEFT JOIN 
		note n ON i.id_scoala = n.id_scoala AND i.id_clasa = n.id_clasa AND i.nume_disciplina = n.nume_disciplina
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

	// Slice-uri pentru a stoca datele necesare pentru grafic
	var xValues []float64
	var yValues []float64
	var textValues [][]string

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
			textValues = append(textValues, []string{fmt.Sprintf("%s: %.2f%%", numeProfesor, procentFeedbackPozitiv.Float64)})
		}
	}

	// Definim structura de date pentru Plotly
	trace1 := ProfessorFeedbackStats{
		X:       xValues,
		Y:       yValues,
		Text:    textValues,
		Mode:    "markers+text", // Setăm modul de afișare pentru puncte și text la hover
		Type:    "scatter",
		Name:    "Professors",
		Marker:  Marker{Size: 12},
		ShowLeg: false, // Opțional: setează true pentru a afișa legenda
	}

	data := []ProfessorFeedbackStats{trace1}

	// Definim layout-ul pentru graficul de tip scatter
	layout := gin.H{
		"xaxis": gin.H{
			"range": [2]interface{}{0.5, 100}, // Interval pentru axa x
		},
		"yaxis": gin.H{
			"range": [2]interface{}{0, 10}, // Interval pentru axa y
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
		"title": "Feedback Positiv vs. Medie Clase",
	}

	// Returnăm datele sub formă de răspuns JSON
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
