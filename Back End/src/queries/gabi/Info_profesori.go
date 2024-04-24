package gabi

import (
	"backend/database"
	"backend/queries/stefan"
	"database/sql"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GenerateToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var token strings.Builder

	for i := 0; i < 10; i++ {
		token.WriteByte(charset[rand.Intn(len(charset))])
	}
	return token.String()
}

func Info_profesor(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := stefan.IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	// Extrage ID-ul școlii din parametrii cererii
	idScoala := context.PostForm("id_scoala")
	if (!stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}

	// Extrage fișierul CSV din corpul cererii
	file, _, err := context.Request.FormFile("csv_file")
	if err != nil {
		fmt.Println("Eroare la extragerea fișierului CSV:", err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Nu s-a putut extrage fișierul CSV"})
		return
	}
	defer file.Close()

	// Deschide fișierul CSV pentru citire
	csvData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Eroare la citirea datelor din fișierul CSV:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la citirea fișierului CSV"})
		return
	}

	// Parcurge fiecare linie din fișierul CSV și inserează profesorul în baza de date
	for _, line := range csvData {
		// Extrage numele și prenumele profesorului din linie
		nume := line[0]    // Presupunând că numele se află pe prima poziție în fiecare linie
		prenume := line[1] // Presupunând că prenumele se află pe a doua poziție în fiecare linie

		// Generează un token aleatoriu de lungime 10
		token := GenerateToken()

		// Generează un ID unic pentru profesor folosind NVL(max(ID), 0) + 1
		var idProfesor int
		query := "SELECT IFNULL(MAX(ID), 0) + 1 FROM profesor WHERE id_scoala = ?"
		err := db.QueryRow(query, idScoala).Scan(&idProfesor)
		if err != nil {
			fmt.Println("Eroare la generarea ID-ului profesorului:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la generarea ID-ului profesorului"})
			return
		}

		// Inserează profesorul în baza de date
		insertStatement := "INSERT INTO profesor (id_scoala, id, nume, prenume, token) VALUES (?, ?, ?, ?, ?)"
		_, err = db.Exec(insertStatement, idScoala, idProfesor, nume, prenume, token)
		if err != nil {
			fmt.Println("Eroare la inserarea profesorului în baza de date:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la inserarea profesorului în baza de date"})
			return
		}
	}

	// Returnează un răspuns de succes
	context.IndentedJSON(http.StatusOK, gin.H{"success": true})
	database.CloseDB(db)

}
