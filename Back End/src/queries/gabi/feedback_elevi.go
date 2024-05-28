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

func Feedback(context *gin.Context) {
	var db *sql.DB = database.InitDb()

	// Extrage continutul feedbackului din corpul cererii
	continut := context.PostForm("Continut feedback")

	// Verifică dacă continutul feedbackului este gol
	if continut == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Continutul feedbackului lipsește"})
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

	// Obține ID-ul profesorului și ID-ul școlii
	var idScoala int
	query := "SELECT id_scoala FROM profesor WHERE id_cont = ?"
	err = db.QueryRow(query, sessionData.ID).Scan(&idScoala)
	if err != nil {
		fmt.Println("Eroare la obținerea ID-ului școlii profesorului:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea ID-ului școlii profesorului"})
		return
	}

	// Generează un ID unic pentru feedback
	var idFeedback int
	query = "SELECT IFNULL(MAX(id_feedback), 0) + 1 FROM feedback WHERE id_scoala = ?"
	err = db.QueryRow(query, idScoala).Scan(&idFeedback)
	if err != nil {
		fmt.Println("Eroare la generarea ID-ului feedbackului:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la generarea ID-ului feedbackului"})
		return
	}

	// Adaugă înregistrarea în tabela "feedback"
	insertFeedbackStatement := "INSERT INTO feedback (id_feedback, id_scoala, nume_disciplina, id_clasa, id_elev, content, data) VALUES (?, ?, ?, ?, ?, ?, NOW())"
	_, err = db.Exec(insertFeedbackStatement, idFeedback, idScoala, context.PostForm("nume_disciplina"), context.PostForm("id_clasa"), context.PostForm("id_elev"), continut)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela feedback:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Eroare la adăugarea înregistrării în tabela feedback"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true})

	database.CloseDB(db)
}
