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

func EleviProfesor(context *gin.Context) {
	var db *sql.DB = database.InitDb()

	// Verificare sesiune activă
	ver := stefan.IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Extrage ID-ul profesorului și al școlii din context
	idProfesor := ver
	idScoala := context.PostForm("id_scoala")

	if idScoala == "" {
		fmt.Println("ID scoala lipseste")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID profesor sau ID scoala lipseste"})
		return
	}

	// Verificare rol profesor pentru școală
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	}) {
		fmt.Println("Userul nu este profesor pentru aceasta scoala")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor pentru aceasta scoala"})
		return
	}

	// Interogare pentru a obține informațiile despre elevi și situația lor școlară
	query := `
		SELECT e.nume, e.prenume, c.nume AS clasa, d.nume_disciplina, n.nota
		FROM elev e
		JOIN clasa c ON e.id_clasa = c.id
		JOIN note n ON e.id = n.id_elev
		JOIN disciplina d ON n.id_disciplina = d.id
		JOIN profesor p ON d.id_profesor = p.id
		WHERE p.id = ? AND p.id_scoala = ?
	`
	rows, err := db.Query(query, idProfesor, idScoala)
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
		Clasa      string `json:"clasa"`
		Disciplina string `json:"disciplina"`
		Nota       int    `json:"nota"`
	}

	var rezultate []SituatieElev

	// Iterarea rezultatelor și adăugarea acestora în structura de rezultate
	for rows.Next() {
		var nume, prenume, clasa, disciplina string
		var nota int
		err := rows.Scan(&nume, &prenume, &clasa, &disciplina, &nota)
		if err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la scanarea rezultatelor"})
			return
		}
		rezultate = append(rezultate, SituatieElev{Nume: nume, Prenume: prenume, Clasa: clasa, Disciplina: disciplina, Nota: nota})
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
