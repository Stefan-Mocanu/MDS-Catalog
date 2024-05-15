package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AdaugaAdmin(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din POST
	idScoala := c.PostForm("id_scoala")
	idCont := c.PostForm("id_cont")
	//Verificare daca useul este PARINTE
	if (!VerificareRol(Rol{
		ROL:    "Parinte",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este parinte in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este parinte in aceasta scoala"})
		return
	}

	q := `insert into cont_rol(id_cont,id_rol,id_scoala)
		values(?,"Admin",?)`
	_, err := db.Exec(q, idCont, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	database.CloseDB(db)
}
