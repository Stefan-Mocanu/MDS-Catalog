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

	// Interogare pentru a obține toate etniile distincte din feedback-uri pentru școala specificată
	q := `
		SELECT DISTINCT etnie
		FROM elev
		WHERE id_scoala = ?
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Slice pentru a stoca etniile
	var etnii []string

	for rows.Next() {
		var ethnicity string
		if err := rows.Scan(&ethnicity); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		etnii = append(etnii, ethnicity)
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
			Label: etnii,                      // Utilizăm aici toate etniile obținute din feedback-uri
			Color: make([]string, len(etnii)), // Alocăm culori pentru fiecare etnie
		},
		Link: Link{
			Source: []int{0, 1, 2, 3, 4, 5},       // Exemplu: Sursa poate fi indexată conform ordinei etniilor
			Target: []int{2, 3, 4, 5, 0, 1},       // Exemplu: Ținta poate fi indexată conform ordinei etniilor
			Value:  []int{10, 20, 30, 40, 50, 60}, // Exemplu: Valoare pentru conexiuni între etnii
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
