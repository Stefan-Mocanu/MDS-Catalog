package gabi

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Funcția pentru adăugarea unei note
func AdaugaActivitate(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	// Extrage valoarea notei din corpul cererii
	valoareNotaStr := c.PostForm("valoare")
	if valoareNotaStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea notei lipsește"})
		return
	}

	// Conversia notei la int
	valoareNota, err := strconv.Atoi(valoareNotaStr)
	if err != nil || valoareNota < 1 || valoareNota > 10 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea notei este invalidă"})
		return
	}

	// Obține cookie-ul de sesiune și validează-l
	cookie, err := c.Cookie("session_cookie")
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"success": false, "error": "Sesiunea nu a fost găsită"})
		return
	}

	// Obține ID-ul profesorului asociat contului din sesiune și ID-ul școlii curente
	var idProfesor, idScoala int
	query := `
		SELECT p.id_profesor, p.id_scoala
		FROM profesor p
		JOIN cont c ON p.id_cont = c.id_cont
		WHERE c.cookie = ?
	`
	err = db.QueryRow(query, cookie).Scan(&idProfesor, &idScoala)
	if err != nil {
		fmt.Println("Eroare la obținerea ID-ului profesorului:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la obținerea ID-ului profesorului"})
		return
	}

	// Verifică dacă profesorul este asociat școlii curente
	idScoalaCurentaStr := c.PostForm("id_scoala")
	if idScoalaCurentaStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID-ul școlii curente lipsește"})
		return
	}
	idScoalaCurenta, err := strconv.Atoi(idScoalaCurentaStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID-ul școlii curente este invalid"})
		return
	}
	if idScoala != idScoalaCurenta {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Profesorul nu este asociat acestei școli"})
		return
	}

	// Adaugă înregistrarea în tabela "note"
	insertNoteStatement := `
		INSERT INTO note (id_scoala, id_profesor, nume_disciplina, id_clasa, id_elev, nota, data)
		VALUES (?, ?, ?, ?, ?, ?, NOW())
	`
	_, err = db.Exec(insertNoteStatement, idScoala, idProfesor, c.PostForm("nume_disciplina"), c.PostForm("id_clasa"), c.PostForm("id_elev"), valoareNota)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela note:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Eroare la adăugarea înregistrării în tabela note"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true})
}
