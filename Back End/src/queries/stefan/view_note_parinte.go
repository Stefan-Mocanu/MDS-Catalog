package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func View_note_parinte(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.Query("id_scoala")
	idClasa := c.Query("id_clasa")
	idElev := c.Query("id_elev")
	if (!VerificareRol(Rol{
		ROL:    "Parinte",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este parinte in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este parinte in aceasta scoala"})
		return
	}
	q := `select 1
		from elev
		where id_scoala = ?
		and id_clasa = ?
		and id_elev = ?
		and id_cont_parinte = ?`
	cnt := 0
	err := db.QueryRow(q, idScoala, idClasa, idElev, ver).Scan(&cnt)
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
