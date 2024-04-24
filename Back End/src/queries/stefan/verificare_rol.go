package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func VerificareRol(deVerificat Rol) bool {
	var db *sql.DB = database.InitDb()
	var rez int
	q := "select count(*) from cont_rol where id_rol = ? and id_cont=? and id-scoala=?"
	err1 := db.QueryRow(q, deVerificat.ROL, deVerificat.ID, deVerificat.SCOALA).Scan(&rez)
	if err1 != nil {
		fmt.Println("Eroare: ", err1)
		return false
	}

	database.CloseDB(db)
	if rez == 1 {
		return true
	} else {
		return false
	}
}
