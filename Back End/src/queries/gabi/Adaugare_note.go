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

func Note(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db) // Asigură închiderea conexiunii la baza de date

	// Extrage valoarea notei din corpul cererii
	valoareNotaStr := context.PostForm("valoare")
	if valoareNotaStr == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea notei lipsește"})
		return
	}

	// Conversia notei la int
	valoareNota, err := strconv.Atoi(valoareNotaStr)
	if err != nil || valoareNota < 1 || valoareNota > 10 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea notei este invalidă"})
		return
	}

	// Obține cookie-ul de sesiune și validează-l
	cookie, err := context.Cookie("session_cookie")
	if err != nil {
		context.IndentedJSON(http.StatusOK, gin.H{"success": false, "error": "Sesiunea nu a fost găsită"})
		return
	}
	sessionData, ok := stefan.Sessions[cookie]
	if !ok {
		context.IndentedJSON(http.StatusOK, gin.H{"success": false, "error": "Sesiunea nu este validă"})
		return
	}

	// Obține ID-ul profesorului din sesiunea activă
	idProfesor := sessionData.ID

	// Extrage ID-ul școlii din formular
	idScoalaStr := context.PostForm("id_scoala")
	if idScoalaStr == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID-ul școlii lipsește"})
		return
	}
	idScoala, err := strconv.Atoi(idScoalaStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID-ul școlii este invalid"})
		return
	}

	// Adaugă înregistrarea în tabela "note"
	insertNoteStatement := "INSERT INTO note (id_scoala, nume_disciplina, id_clasa, id_elev, nota, data, id_profesor) VALUES (?, ?, ?, ?, ?, NOW(), ?)"
	_, err = db.Exec(insertNoteStatement, idScoala, context.PostForm("nume_disciplina"), context.PostForm("id_clasa"), context.PostForm("id_elev"), valoareNota, idProfesor)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela note:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Eroare la adăugarea înregistrării în tabela note"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
