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

func Info_incadrare(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	// Extrage ID-ul școlii din parametrii cererii
	idScoala := context.PostForm("id_scoala")
	if (!VerificareRol(Rol{
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
	prof := make(map[string]int)
	q := `select nume ||" "||prenume, id from profesor where id_scoala = ?;`
	rows, err1 := db.Query(q, idScoala)
	if err1 != nil {
		fmt.Println("Eroare: ", err1)
		context.IndentedJSON(http.StatusOK, false)
		return
	}
	for rows.Next() {
		var aux1 string
		var aux2 int
		if err := rows.Scan(&aux1, &aux2); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			prof[aux1] = aux2
		}
	}

	// Deschide fișierul CSV pentru citire
	csvData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Eroare la citirea datelor din fișierul CSV:", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Eroare la citirea fișierului CSV"})
		return
	}

	for _, line := range csvData {

		clasa := line[0]
		nume := line[1]
		materie := line[2]

		// Inserează incadrarea în baza de date
		insertStatement := "INSERT INTO incadreare (id_scoala, id_clasa, id_profesor,nume_disciplina) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(insertStatement, idScoala, clasa, prof[nume], materie)
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
