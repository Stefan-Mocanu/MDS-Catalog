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

func Prezenta(context *gin.Context) {
	var db *sql.DB = database.InitDb()

	// Extrage valoarea activității din corpul cererii
	valoareActivitateStr := context.PostForm("valoare")
	if valoareActivitateStr == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea activității lipsește"})
		return
	}

	// Conversia valorii activității la int
	valoareActivitate, err := strconv.Atoi(valoareActivitateStr)
	if err != nil || (valoareActivitate != 0 && valoareActivitate != 1) {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Valoarea activității este invalidă, trebuie să fie 0 sau 1"})
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

	// Generează un ID unic pentru activitate
	var idActivitate int
	query = "SELECT IFNULL(MAX(id_nota), 0) + 1 FROM activitate WHERE id_scoala = ?"
	err = db.QueryRow(query, idScoala).Scan(&idActivitate)
	if err != nil {
		fmt.Println("Eroare la generarea ID-ului activității:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la generarea ID-ului activității"})
		return
	}

	// Adaugă înregistrarea în tabela "activitate"
	insertActivitateStatement := "INSERT INTO activitate (id_nota, id_scoala, nume_disciplina, id_clasa, id_elev, valoare, data) VALUES (?, ?, ?, ?, ?, ?, NOW())"
	_, err = db.Exec(insertActivitateStatement, idActivitate, idScoala, context.PostForm("nume_disciplina"), context.PostForm("id_clasa"), context.PostForm("id_elev"), valoareActivitate)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela activitate:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Eroare la adăugarea înregistrării în tabela activitate"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"success": true})

	database.CloseDB(db)
}
