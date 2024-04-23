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

func Inregistrare_Instit_Invat(context *gin.Context) {
	var db *sql.DB = database.InitDb()

	// Extrage numele școlii din corpul cererii
	nume_Scoala := context.PostForm("nume")

	// Verifică dacă numele școlii este gol
	if nume_Scoala == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Numele școlii lipsește"})
		return
	}

	insertStatement := "INSERT INTO scoala (nume) VALUES (?)"

	result, err := db.Exec(insertStatement, nume_Scoala)
	if err != nil {
		fmt.Println("Eroare la adăugarea instituției de învățământ:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// Obține ID-ul instituției de învățământ adăugate
	newSchoolID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Eroare la obținerea ID-ului instituției de învățământ adăugate:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// Adaugă înregistrarea în tabela "cont_rol"
	cookie, err := context.Cookie("session_cookie")
	if err != nil {
		context.IndentedJSON(http.StatusOK, false)
		return
	}
	content, ok := stefan.Sessions[cookie]
	if !ok {
		context.IndentedJSON(http.StatusOK, false)
		return
	}

	insertContRolStatement := "INSERT INTO cont_rol (id_cont, id_rol, id_scoala) VALUES (?, ?, ?)"
	_, err = db.Exec(insertContRolStatement, content.ID, "Administrator", newSchoolID)
	if err != nil {
		fmt.Println("Eroare la adăugarea înregistrării în tabela cont_rol:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// Returnează răspunsul HTTP de succes și ID-ul instituției de învățământ adăugate
	context.IndentedJSON(http.StatusOK, gin.H{"success": true, "id": newSchoolID})

	// Închide conexiunea la baza de date
	database.CloseDB(db)

}
