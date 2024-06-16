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

type EvolutieNota struct {
	X    []int  `json:"x"`
	Y    []int  `json:"y"`
	NUME string `json:"name"`
	TIP  string `json:"type"`
}

func GetEvolutieNoteElevi(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Ensure database connection closure

	// Check if session is active
	ver := stefan.IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extract class ID and school ID from query
	idClasa := c.Query("id_clasa")
	if idClasa == "" {
		fmt.Println("ID clasa lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID clasa lipseste"})
		return
	}

	idScoalaStr := c.Query("id_scoala")
	if idScoalaStr == "" {
		fmt.Println("ID scoala lipseste")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipseste"})
		return
	}
	idScoala, err := strconv.Atoi(idScoalaStr)
	if err != nil {
		fmt.Println("ID scoala este invalid")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID scoala este invalid"})
		return
	}

	// Check if user has 'Profesor' role for the school
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Profesor",
		SCOALA: strconv.Itoa(idScoala),
		ID:     ver,
	}) {
		fmt.Println("Userul nu este profesor pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor pentru aceasta scoala"})
		return
	}

	// Query to fetch students' grades for the specified class
	query := `
		SELECT YEAR(n.data), n.nota, e.id_elev, e.nume
		FROM note n
		JOIN elev e ON n.id_elev = e.id_elev AND n.id_scoala = e.id_scoala AND n.id_clasa = e.id_clasa
		WHERE e.id_clasa = ? AND e.id_scoala = ?
		ORDER BY e.id_elev, n.data
	`
	rows, err := db.Query(query, idClasa, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Structures to store results
	eleviNote := map[int][]EvolutieNota{}
	var ani []int

	for rows.Next() {
		var an, nota, elevID int
		var numeElev string
		err := rows.Scan(&an, &nota, &elevID, &numeElev)
		if err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la scanarea rezultatelor"})
			return
		}

		// Store data in map and unique years in slice
		eleviNote[elevID] = append(eleviNote[elevID], EvolutieNota{X: []int{an}, Y: []int{nota}, NUME: numeElev})
		if !contains(ani, an) {
			ani = append(ani, an)
		}
	}

	if len(eleviNote) == 0 {
		fmt.Println("Nu s-au găsit note pentru elevii din această clasă.")
		c.IndentedJSON(http.StatusOK, gin.H{"data": nil, "layout": map[string]interface{}{
			"barmode": "stack",
		}})
		return
	}

	// Prepare data for JSON response
	var data []EvolutieNota
	for _, note := range eleviNote {
		var x []int
		var y []int
		var numeElev string
		for _, nota := range note {
			x = append(x, nota.X...)
			y = append(y, nota.Y...)
			numeElev = nota.NUME
		}
		data = append(data, EvolutieNota{
			X:    x,
			Y:    y,
			NUME: numeElev,
			TIP:  "scatter",
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data, "layout": map[string]interface{}{
		"barmode": "stack",
	}})
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
