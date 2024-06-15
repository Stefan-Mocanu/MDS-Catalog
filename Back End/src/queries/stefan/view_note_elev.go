package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Note struct {
	DATA string `json:"data"`
	NOTA int    `json:"nota"`
}
type Absente struct {
	DATA    string `json:"data"`
	MATERIE string `json:"materie"`
}
type Feedback struct {
	CONTENT string `json:"content"`
	DATA    string `json:"data"`
	TIP     bool   `json:"tip"`
}

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func View_note_elev(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.Query("id_scoala")
	idClasa := c.Query("id_clasa")
	if (!VerificareRol(Rol{
		ROL:    "Elev",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este elev in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este elev in aceasta scoala"})
		return
	}
	q := `select id_elev
		from elev
		where id_scoala = ?
		and id_clasa = ?
		and id_cont_elev = ?`
	idElev := 0
	err := db.QueryRow(q, idScoala, idClasa, ver).Scan(&idElev)
	switch {

	case err == sql.ErrNoRows:
		fmt.Printf("Eroare: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Nu exista acest elev."})
	case err != nil:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
	}
	q = `select nume_disciplina, nota, data
		from note
		where id_scoala = ?
		and id_clasa = ?
		and id_elev = ?
		order by data`
	rows, err := db.Query(q, idScoala, idClasa, idElev)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	var catalog_note = map[string][]Note{}
	var catalog_activitate = map[string][]Note{}
	var catalog_feedback = map[string][]Feedback{}
	var absente = []Absente{}
	for rows.Next() {
		var materie, data string
		var nota int
		if err := rows.Scan(&materie, &nota, &data); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			catalog_note[materie] = append(catalog_note[materie], Note{
				NOTA: nota,
				DATA: data,
			})
		}
	}
	q = `select nume_disciplina, valoare, data
		from activitate
		where id_scoala = ?
		and id_clasa = ?
		and id_elev = ?
		order by data`
	rows, err = db.Query(q, idScoala, idClasa, idElev)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var materie, data string
		var nota int
		if err := rows.Scan(&materie, &nota, &data); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			if nota == 0 {
				absente = append(absente, Absente{
					DATA:    data,
					MATERIE: materie,
				})
				continue
			}
			catalog_activitate[materie] = append(catalog_activitate[materie], Note{
				NOTA: nota,
				DATA: data,
			})
		}
	}
	q = `select nume_disciplina, content, data, tip
		from feedback
		where id_scoala = ?
		and id_clasa = ?
		and id_elev = ?
		and directie = 0
		order by data`
	rows, err = db.Query(q, idScoala, idClasa, idElev)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var materie, data, content string
		var tip bool
		if err := rows.Scan(&materie, &content, &data, &tip); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			catalog_feedback[materie] = append(catalog_feedback[materie], Feedback{
				CONTENT: content,
				DATA:    data,
				TIP:     tip,
			})
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"Note":       catalog_note,
		"Activitate": catalog_activitate,
		"Feedback":   catalog_feedback,
		"Absente":    absente,
	})

	database.CloseDB(db)
}
