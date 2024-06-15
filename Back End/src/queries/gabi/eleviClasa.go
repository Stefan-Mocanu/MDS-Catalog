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

func EleviClasa(context *gin.Context) {
	var db *sql.DB = database.InitDb()

	// Verificare sesiune activă
	ver := stefan.IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extrage ID-ul clasei din context
	idClasa := context.PostForm("id_clasa")
	if idClasa == "" {
		fmt.Println("ID clasa lipseste")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID clasa lipseste"})
		return
	}

	// Verificare rol profesor pentru școală
	idScoala := context.PostForm("id_scoala")
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este profesor pentru aceasta scoala")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor pentru aceasta scoala"})
		return
	}

	// Interogare pentru a obține situația școlară a elevilor din clasa specificată
	query := `
		SELECT e.nume, e.prenume, d.nume, n.nota
		FROM elev e
		JOIN note n ON e.id_elev = n.id_elev and e.id_scoala = n.id_scoala and e.id_clasa = n.id_clasa
		JOIN discipline d ON n.nume_disciplina = d.nume
		WHERE e.id_clasa = ?
	`
	rows, err := db.Query(query, idClasa)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Structură pentru rezultatele interogării
	type SituatieElev struct {
		Nume       string `json:"nume"`
		Prenume    string `json:"prenume"`
		Disciplina string `json:"disciplina"`
		Nota       int    `json:"nota"`
	}

	var rezultate []SituatieElev

	// Iterarea rezultatelor și adăugarea acestora în structura de rezultate
	for rows.Next() {
		var nume, prenume, disciplina string
		var nota int
		err := rows.Scan(&nume, &prenume, &disciplina, &nota)
		if err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la scanarea rezultatelor"})
			return
		}
		rezultate = append(rezultate, SituatieElev{Nume: nume, Prenume: prenume, Disciplina: disciplina, Nota: nota})
	}

	// Verificare erori în timpul iterării
	if err = rows.Err(); err != nil {
		fmt.Println("Eroare la iterarea rezultatelor:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la iterarea rezultatelor"})
		return
	}

	// Returnarea rezultatelor în format JSON
	context.IndentedJSON(http.StatusOK, gin.H{"data": rezultate})

	database.CloseDB(db)

}
