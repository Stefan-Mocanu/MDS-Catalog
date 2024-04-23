package stefan

import (
	"backend/database"
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Info_Materii(context *gin.Context) {
	// Inițializează conexiunea la baza de date
	var db *sql.DB = database.InitDb()

	// Extrage ID-ul școlii din parametrii cererii
	idScoala := context.PostForm("id_scoala")

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

	// Parcurge fiecare linie din fișierul CSV și inserează clasa în baza de date
	for _, line := range csvData {
		// Extrage numele clasei din linie
		numeMaterie := line[0] // Presupunând că numele clasei se află pe prima poziție în fiecare linie

		// Inserează clasa în baza de date
		insertStatement := "INSERT INTO disciplina (id_scoala, nume) VALUES (?, ?)"
		_, err := db.Exec(insertStatement, idScoala, numeMaterie)
		if err != nil {
			fmt.Println("Eroare la inserarea clasei în baza de date:", err)
			// Poți să decizi dacă vrei să continui sau să oprești procesarea în caz de eroare
		}
	}

	// Returnează un răspuns de succes
	context.IndentedJSON(http.StatusOK, gin.H{"success": true})

	database.CloseDB(db)

}
