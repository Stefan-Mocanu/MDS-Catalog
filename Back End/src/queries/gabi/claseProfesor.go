package gabi

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Clase_Profesor(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	// Verificare sesiune activă și obținere id_cont din sesiune
	cookie, err := context.Cookie("session_cookie")
	if err != nil {
		fmt.Println("Sesiunea nu a fost găsită")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Sesiunea nu a fost găsită"})
		return
	}

	// Obținerea id_cont din sesiune
	var idCont int
	query := "SELECT id_cont FROM cont WHERE cookie = ?"
	err = db.QueryRow(query, cookie).Scan(&idCont)
	if err != nil {
		fmt.Println("Eroare la obținerea id_cont din sesiune:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea id_cont din sesiune"})
		return
	}

	// Obținerea id_profesor asociat id_cont din sesiune
	var idProfesor int
	query = "SELECT id_profesor FROM profesor WHERE id_cont = ?"
	err = db.QueryRow(query, idCont).Scan(&idProfesor)
	if err != nil {
		fmt.Println("Eroare la obținerea id_profesor din baza de date:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea id_profesor din baza de date"})
		return
	}

	// Verificare rol profesor pentru școală
	idScoala := context.PostForm("id_scoala")
	if !VerificaRolProfesor(idProfesor, idScoala, db) {
		fmt.Println("Profesorul nu este asociat acestei școli sau rolul nu este Profesor")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Profesorul nu este asociat acestei școli sau rolul nu este Profesor"})
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

// Funcție pentru verificarea rolului profesorului pentru o școală specificată
func VerificaRolProfesor(idProfesor int, idScoala string, db *sql.DB) bool {
	query := `
		SELECT COUNT(*)
		FROM profesor
		WHERE id_profesor = ? AND id_scoala = ? AND rol = "Profesor"
	`
	var count int
	err := db.QueryRow(query, idProfesor, idScoala).Scan(&count)
	if err != nil {
		fmt.Println("Eroare la verificarea rolului profesorului:", err)
		return false
	}
	return count > 0
}
