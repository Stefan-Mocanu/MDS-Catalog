package stefan

import (
	"backend/database"
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateCSVprofesor(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	// Extrage ID-ul școlii din parametrii cererii
	idScoala := c.Query("id_scoala")
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Verificare daca userul este admin
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}
	records := [][]string{}
	headers := []string{"Nume", "Token"}
	records = append([][]string{headers}, records...)
	//Obtinere date token si date de identificare a profesorilor
	q := `select concat(nume," ",prenume), token from profesor where id_scoala = ? order by nume, prenume`
	rows, err1 := db.Query(q, idScoala)
	if err1 != nil {
		fmt.Println("Eroare: ", err1)
		c.IndentedJSON(http.StatusOK, false)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var record = []string{"", ""}
		if err := rows.Scan(&record[0], &record[1]); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			records = append(records, record)
		}
	}
	if len(records) == 0 {
		c.IndentedJSON(http.StatusOK, false)
		return
	}
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write data rows to CSV buffer
	for _, row := range records {
		if err := writer.Write(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	writer.Flush()

	// Set MIME type explicitly to CSV
	c.Header("Content-Type", "text/csv")

	// Serve CSV file from buffer
	c.Data(http.StatusOK, "text/csv", buf.Bytes())

}
