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

	// Obține ID-ul profesorului și ID-ul școlii
	var idScoala int
	query := "SELECT id_scoala FROM profesor WHERE id_cont = ?"
	err = db.QueryRow(query, sessionData.ID).Scan(&idScoala)
	if err != nil {
		fmt.Println("Eroare la obținerea ID-ului școlii profesorului:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea ID-ului școlii profesorului"})
		return
	}

	// Generează un ID unic pentru nota
	var idNota int
	query = "SELECT IFNULL(MAX(id_nota), 0) + 1 FROM note WHERE id_scoala = ?"
	err = db.QueryRow(query, idScoala).Scan(&idNota)
	if err != nil {
		fmt.Println("Eroare la generarea ID-ului notei:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la generarea ID-ului notei"})
		return
	}

	// Adaugă înregistrarea în tabela "note"
	insertNoteStatement := "INSERT INTO note (id_nota, id_scoala, nume_disciplina, id_clasa, id_elev, nota, data) VALUES (?, ?, ?, ?, ?, ?, NOW())"
	_, err = db.Exec(insertNoteStatement, idNota, idScoala, context.PostForm("nume_disciplina"), context.PostForm("id_clasa"), context.PostForm("id_elev"), valoareNota)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela note:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Eroare la adăugarea înregistrării în tabela note"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true})

	database.CloseDB(db)
}
