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

	// Extract school ID from query string
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

	// Verify user has Administrator role for the specified school
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoalaStr,
		ID:     ver,
	}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin în această școală"})
		return
	}

	// Query to get the number of positive and negative feedbacks per ethnicity
	q := `
		SELECT e.etnie, 
		       SUM(CASE WHEN f.tip = 1 THEN 1 ELSE 0 END) AS pozitive,
		       SUM(CASE WHEN f.tip = 0 THEN 1 ELSE 0 END) AS negative
		FROM elev e
		JOIN feedback f ON e.id_elev = f.id_elev
		WHERE e.id_scoala = ?
		GROUP BY e.etnie
	`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	var etnii []string
	var pozitive []int
	var negative []int

	for rows.Next() {
		var etnie string
		var poz int
		var neg int
		if err := rows.Scan(&etnie, &poz, &neg); err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			continue
		}
		etnii = append(etnii, etnie)
		pozitive = append(pozitive, poz)
		negative = append(negative, neg)
	}

	// Define Sankey data structure for Plotly
	var labels = []string{}
	var colors = []string{}
	var sources = []int{}
	var targets = []int{}
	var values = []int{}

	// Add nodes for each ethnicity
	for i, etnie := range etnii {
		labels = append(labels, etnie+" Pozitiv")
		labels = append(labels, etnie+" Negativ")
		colors = append(colors, "green")
		colors = append(colors, "red")
		sources = append(sources, i*2)            // Pozitiv
		sources = append(sources, i*2+1)          // Negativ
		targets = append(targets, len(etnii)*2)   // Target for Pozitiv
		targets = append(targets, len(etnii)*2+1) // Target for Negativ
		values = append(values, pozitive[i])
		values = append(values, negative[i])
	}

	// Add nodes for total positive and negative feedbacks
	labels = append(labels, "Total Pozitiv")
	labels = append(labels, "Total Negativ")
	colors = append(colors, "blue")
	colors = append(colors, "blue")

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
			Label: labels,
			Color: colors,
		},
		Link: Link{
			Source: sources,
			Target: targets,
			Value:  values,
		},
	}

	data := []SankeyData{sankeyData}

	// Define layout for Sankey chart
	layout := gin.H{
		"title": "Feedback Pozitive/Negative pe Etnie",
		"font": gin.H{
			"size": 10,
		},
	}

	// Return JSON response
	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": layout})
}
