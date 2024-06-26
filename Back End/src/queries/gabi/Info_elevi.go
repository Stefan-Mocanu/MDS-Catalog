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
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GenerateUniqueToken1(db *sql.DB) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var token strings.Builder

	for {
		for i := 0; i < 10; i++ {
			token.WriteByte(charset[r.Intn(len(charset))])
		}
		newToken := token.String()

		// Verifică unicitatea tokenului în baza de date
		var count int
		query := "SELECT COUNT(*) FROM (SELECT token_elev AS token FROM elev UNION SELECT token_parinte AS token FROM elev UNION SELECT token FROM profesor) AS all_tokens WHERE token = ?"
		err := db.QueryRow(query, newToken).Scan(&count)
		if err != nil {
			fmt.Println("Eroare la verificarea unicității tokenului:", err)
			continue
		}

		if count == 0 {
			return newToken
		}

		token.Reset()
	}
}

func Info_elev(context *gin.Context) {
	var db *sql.DB = database.InitDb()
	defer database.CloseDB(db)

	ver := stefan.IsSessionActiveIntern(context)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	// Extrage ID-ul școlii din parametrii cererii
	idScoala := context.PostForm("id_scoala")
	if !stefan.VerificareRol(stefan.Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	}) {
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

	// Parcurge fiecare linie din fișierul CSV și inserează elevul în baza de date
	for _, line := range csvData {
		// Extrage datele elevului din linie
		idClasa := line[0]
		nume := line[1]
		prenume := line[2]
		gen := line[3]
		etnie := line[4]

		// Generează un token aleatoriu pentru elev și părinte
		tokenElev := GenerateUniqueToken1(db)
		tokenParinte := GenerateUniqueToken1(db)

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
}
