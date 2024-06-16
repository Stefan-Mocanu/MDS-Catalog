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

type MedieClasa struct {
	Clasa int     `json:"clasa"`
	Medie float64 `json:"medie"`
}

func MediiClase(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Asigură închiderea conexiunii la baza de date

	// Verificăm dacă sesiunea este activă
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extragem id-ul profesorului din query string
	idProfesorStr := c.Query("id_profesor")
	if idProfesorStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID profesor lipsește"})
		return
	}
	idProfesor, err := strconv.Atoi(idProfesorStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID profesor invalid"})
		return
	}

	// Verificăm dacă utilizatorul are rolul de Profesor și este asociat cu profesorul specificat
	if !stefan.VerificareRol(stefan.Rol{
		ROL: "Profesor",
		ID:  ver,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor"})
		return
	}

	// Interogarea pentru a obține media la fiecare clasă la care predă profesorul
	q := `
		SELECT n.id_clasa, AVG(n.nota) AS medie
		FROM note n
		JOIN elev e ON n.id_elev = e.id_elev AND n.id_scoala = e.id_scoala AND n.id_clasa = e.id_clasa
		WHERE e.id_profesor = ?
		GROUP BY n.id_clasa
	`
	rows, err := db.Query(q, idProfesor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice pentru a stoca mediile claselor
	var mediiClase []MedieClasa

	for rows.Next() {
		var medieClasa MedieClasa
		if err := rows.Scan(&medieClasa.Clasa, &medieClasa.Medie); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		mediiClase = append(mediiClase, medieClasa)
	}

	// Datele pentru grafic
	var xValues []int
	var yValues []float64

	for _, medieClasa := range mediiClase {
		xValues = append(xValues, medieClasa.Clasa)
		yValues = append(yValues, medieClasa.Medie)
	}

	// Returnăm datele într-un format JSON compatibil cu Plotly
	data := []map[string]interface{}{
		{
			"type":     "funnel",
			"name":     "Medii Clase",
			"y":        xValues, // Folosim xValues ca și y pentru a afișa mediile pe axe
			"x":        yValues, // Mediile vor fi valorile pentru axa x în grafic
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
