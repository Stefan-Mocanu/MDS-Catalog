package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func AlaturareElev(c *gin.Context) {
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}

	var db *sql.DB = database.InitDb()
	token := c.PostForm("token")
	idScoala := c.PostForm("id_scoala")
	var id int
	q := `SELECT id_elev
	FROM elev
	WHERE id_scoala = ? and token_elev = ?`
	err := db.QueryRow(q, idScoala, token).Scan(&id)
	switch {

	case err == sql.ErrNoRows:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Nu exista elev cu acest cod in baza de date"})
		return
	case err != nil:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}

	q = `update elev
		set id_cont_elev = ?
		where id_scoala = ?
		and token = ?`
	_, err = db.Exec(q, ver, idScoala, token)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	q = `insert into cont_rol(id_cont,id_rol,id_scoala)
		values(?,"Elev",?)`
	_, err = db.Exec(q, ver, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	database.CloseDB(db)
}
