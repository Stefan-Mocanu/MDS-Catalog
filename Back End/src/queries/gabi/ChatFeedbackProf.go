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

type FeedbackCount struct {
	Profesor string `json:"profesor"`
	Numar    int    `json:"numar"`
}

func FeedbackuriProfesori(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	// Verificăm dacă sesiunea este activă
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.Query("id_scoala")
	if idScoala == "" {
		fmt.Println("ID scoala lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipseste"})
		return
	}

	// Verificăm dacă utilizatorul are rolul de Administrator sau alt rol necesar
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		ID:     ver,
		SCOALA: idScoala,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Accesul interzis"})
		return
	}

	// Interogare pentru a obține numărul de feedback-uri pentru fiecare profesor
	q := `
		SELECT p.nume AS profesor, COUNT(f.id_feedback) AS numar
		FROM feedback f
		JOIN incadrare i ON f.id_scoala = i.id_scoala AND f.id_clasa = i.id_clasa AND f.nume_disciplina = i.nume_disciplina
		JOIN profesor p ON i.id_profesor = p.id AND i.id_scoala = p.id_scoala
		WHERE f.id_scoala = ?
		GROUP BY p.nume
		ORDER BY numar DESC
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice pentru a stoca datele necesare pentru graficul de tip bar
	var feedbackCounts []FeedbackCount

	for rows.Next() {
		var profesor string
		var numar int
		if err := rows.Scan(&profesor, &numar); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		feedbackCounts = append(feedbackCounts, FeedbackCount{Profesor: profesor, Numar: numar})
	}

	// Construim datele sub forma necesară pentru JavaScript (JSON)
	var xValues []string
	var yValues []int
	for _, feedback := range feedbackCounts {
		xValues = append(xValues, feedback.Profesor)
		yValues = append(yValues, feedback.Numar)
	}

	// Construim structura de date pentru Plotly
	data := []map[string]interface{}{
		{
			"x":    xValues,
			"y":    yValues,
			"type": "bar",
		},
	}

	// Returnăm datele sub formă de răspuns JSON
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}
