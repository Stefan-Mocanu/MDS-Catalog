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

func HeatMapMediiLunileAnului(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Asigură închiderea conexiunii la baza de date

	// Verificăm dacă sesiunea este activă
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extragem id-ul școlii din query string
	idScoala := c.Query("id_scoala")
	if idScoala == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipsește"})
		return
	}

	// Verificăm dacă utilizatorul are rolul de Administrator pentru școala specificată
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este admin în această scoală")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această scoală"})
		return
	}

	// Definim valorile pentru x (materii) și y (luni)
	xValues := []string{"A", "B", "C", "D", "E"} // Exemplu de materii
	yValues := []string{"W", "X", "Y", "Z"}      // Exemplu de luni

	// Declaram matricea z pentru a stoca datele
	zValues := [][]float64{
		{0.00, 0.00, 0.75, 0.75, 0.00},
		{0.00, 0.00, 0.75, 0.75, 0.00},
		{0.75, 0.75, 0.75, 0.75, 0.75},
		{0.00, 0.00, 0.00, 0.75, 0.00},
	}

	// Interogare pentru a obține mediile pe materii și lunile anului
	q := `
		SELECT 
			materie,
			luna,
			AVG(medie) AS avg_medie
		FROM 
			medii
		WHERE 
			id_scoala = ?
		GROUP BY 
			materie, luna
		ORDER BY 
			materie, luna
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Map pentru a stoca mediile pe materii și lunile anului
	mediile := make(map[string]map[string]float64)

	// Iterăm prin rezultate și populăm map-ul mediile
	for rows.Next() {
		var materie, luna string
		var avgMedie float64
		if err := rows.Scan(&materie, &luna, &avgMedie); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		if _, ok := mediile[materie]; !ok {
			mediile[materie] = make(map[string]float64)
		}
		mediile[materie][luna] = avgMedie
	}

	// Construim zValues bazat pe mediile obținute din baza de date
	for i, materie := range xValues {
		for j, luna := range yValues {
			if medie, ok := mediile[materie][luna]; ok {
				zValues[i][j] = medie
			}
		}
	}

	// Definim scale-ul de culori pentru heatmap
	colorscaleValue := [][]interface{}{
		{0, "#3D9970"},
		{1, "#001f3f"},
	}

	// Construim datele pentru Plotly heatmap
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

	// Definim layout-ul pentru heatmap
	layout := map[string]interface{}{
		"title": "Heatmap Medii/Luni Anului",
		"xaxis": map[string]interface{}{
			"title": "Materii",
		},
		"yaxis": map[string]interface{}{
			"title": "Luni",
		},
	}

	// Returnăm datele sub forma unui răspuns JSON
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
