package gabi

import (
	"backend/database"
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Info_elev(context *gin.Context) {
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

	// Parcurge fiecare linie din fișierul CSV și inserează elevul în baza de date
	for _, line := range csvData {
		// Extrage datele elevului din linie
		idClasa := line[0]
		nume := line[1]
		prenume := line[2]
		gen := line[3]
		etnie := line[4]

		// Generează un token aleatoriu pentru elev și părinte
		tokenElev := GenerateToken()
		tokenParinte := GenerateToken()

		// Generează un ID unic pentru elev folosind NVL(max(ID), 0) + 1
		var idElev int
		query := "SELECT IFNULL(MAX(id_elev), 0) + 1 FROM elev WHERE id_scoala = ? AND id_clasa = ?"
		err := db.QueryRow(query, idScoala, idClasa).Scan(&idElev)
		if err != nil {
			fmt.Println("Eroare la generarea ID-ului elevului:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la generarea ID-ului elevului"})
			return
		}

		// Inserează elevul în baza de date
		insertStatement := "INSERT INTO elev (id_scoala, id_clasa, id_elev, nume, prenume, gen, etnie, token_elev, token_parinte) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(insertStatement, idScoala, idClasa, idElev, nume, prenume, gen, etnie, tokenElev, tokenParinte)
		if err != nil {
			fmt.Println("Eroare la inserarea elevului în baza de date:", err)
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la inserarea elevului în baza de date"})
			return
		}
	}

	// Returnează un răspuns de succes
	context.IndentedJSON(http.StatusOK, gin.H{"success": true})

	database.CloseDB(db)

}
