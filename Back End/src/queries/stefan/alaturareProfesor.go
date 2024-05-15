package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AlaturareProf(c *gin.Context) {
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	var db *sql.DB = database.InitDb()
	//Obtinere date din POST
	token := c.PostForm("token")
	var idScoala int
	//Obtinere id profesor
	q := `SELECT idScoala
	FROM profesor 
	WHERE token = ?`
	err := db.QueryRow(q, token).Scan(&idScoala)
	switch {

	case err == sql.ErrNoRows:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Nu exista profesor cu acest cod in baza de date"})
		return
	case err != nil:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	//Linkuire cont profesor
	q = `update profesor
		set id_cont = ?
		where token = ?`
	_, err = db.Exec(q, ver, token)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	//Adaugare rol in tabelul de roluri
	q = `insert into cont_rol(id_cont,id_rol,id_scoala)
		values(?,"Profesor",?)`
	_, err = db.Exec(q, ver, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	database.CloseDB(db)
}
