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

type SankeyData struct {
	Type        string `json:"type"`
	Orientation string `json:"orientation"`
	Node        Node   `json:"node"`
	Link        Link   `json:"link"`
}

type Node struct {
	Pad       int      `json:"pad"`
	Thickness int      `json:"thickness"`
	Line      Line     `json:"line"`
	Label     []string `json:"label"`
	Color     []string `json:"color"`
}

type Line struct {
	Color string  `json:"color"`
	Width float64 `json:"width"`
}

type Link struct {
	Source []int `json:"source"`
	Target []int `json:"target"`
	Value  []int `json:"value"`
}

func EthnicitySankey(c *gin.Context) {
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

	// Interogare pentru a obține feedback-urile pozitive și negative pentru fiecare etnie
	q := `
		SELECT 
			IF(ethnicity = 'A1', 0, IF(ethnicity = 'A2', 1, IF(ethnicity = 'B1', 2, IF(ethnicity = 'B2', 3, IF(ethnicity = 'C1', 4, 5))))) AS source,
			IF(feedback_pozitiv, IF(ethnicity = 'A1' OR ethnicity = 'A2', 2, IF(ethnicity = 'B1' OR ethnicity = 'B2', 3, 4)), 5) AS target,
			COUNT(*) AS value
		FROM 
			feedback
		WHERE 
			id_scoala = ?
		GROUP BY 
			source, target
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice-uri pentru a stoca datele necesare pentru grafic
	var sourceValues []int
	var targetValues []int
	var valueValues []int

	for rows.Next() {
		var source, target, value int
		if err := rows.Scan(&source, &target, &value); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		sourceValues = append(sourceValues, source)
		targetValues = append(targetValues, target)
		valueValues = append(valueValues, value)
	}

	// Definim structura de date pentru Plotly Sankey
	sankeyData := SankeyData{
		Type:        "sankey",
		Orientation: "h",
		Node: Node{
			Pad:       15,
			Thickness: 30,
			Line: Line{
				Color: "black",
				Width: 0.5,
			},
			Label: []string{"A1", "A2", "B1", "B2", "C1", "C2"},
			Color: []string{"blue", "blue", "blue", "blue", "blue", "blue"},
		},
		Link: Link{
			Source: sourceValues,
			Target: targetValues,
			Value:  valueValues,
		},
	}

	data := []SankeyData{sankeyData}

	// Definim layout-ul pentru graficul Sankey
	layout := gin.H{
		"title": "Feedback Pozitive/Negative pe Etnie",
		"font": gin.H{
			"size": 10,
		},
	}

	// Returnăm datele sub formă de răspuns JSON
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
