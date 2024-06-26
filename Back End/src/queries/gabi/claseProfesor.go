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

func Clase_Profesor(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	// Verificare sesiune activă și obținere ver din sesiune
	ver := stefan.IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	// Obținerea idcont din sesiune
	idcont := ver

	// Obținerea id_scoala din context
	idScoala := context.PostForm("id_scoala")
	if idScoala == "" {
		fmt.Println("ID scoala lipseste")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID scoala lipseste"})
		return
	}

	// Interogare pentru a obține id_profesor din baza de date
	var idProfesor int
	query := "SELECT id FROM profesor WHERE id_cont = ? AND id_scoala = ?"
	err := db.QueryRow(query, idcont, idScoala).Scan(&idProfesor)
	if err != nil {
		fmt.Println("Eroare la obținerea id_profesor din baza de date:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea id_profesor din baza de date"})
		return
	}

	// Verificare rol profesor pentru școală
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     idcont,
	}) {
		fmt.Println("Userul nu este profesor pentru aceasta scoala")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor pentru aceasta scoala"})
		return
	}

	// Interogare pentru a obține clasele și disciplinele profesorului
	query = `
		SELECT id_clasa, nume_disciplina
		FROM incadrare
		WHERE id_profesor = ? AND id_scoala = ?
	`
	rows, err := db.Query(query, idProfesor, idScoala)
	if err != nil {
		fmt.Println("Eroare la interogarea bazei de date:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la interogarea bazei de date"})
		return
	}
	defer rows.Close()

	// Structură pentru rezultatele interogării
	type ClasaDisciplina struct {
		Clasa      string `json:"clasa"`
		Disciplina string `json:"disciplina"`
	}

	var rezultate []ClasaDisciplina

	// Iterarea rezultatelor și adăugarea acestora în structura de rezultate
	for rows.Next() {
		var clasa, disciplina string
		err := rows.Scan(&clasa, &disciplina)
		if err != nil {
			fmt.Println("Eroare la scanarea rezultatelor:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la scanarea rezultatelor"})
			return
		}
		rezultate = append(rezultate, ClasaDisciplina{Clasa: clasa, Disciplina: disciplina})
	}

	// Verificare erori în timpul iterării
	if err = rows.Err(); err != nil {
		fmt.Println("Eroare la iterarea rezultatelor:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la iterarea rezultatelor"})
		return
	}

	// Returnarea rezultatelor în format JSON
	context.IndentedJSON(http.StatusOK, gin.H{"data": rezultate})
}
