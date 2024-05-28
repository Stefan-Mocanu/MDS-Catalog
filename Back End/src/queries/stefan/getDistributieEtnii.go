package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	X    []string `json:"x"`
	Y    []int    `json:"y"`
	NUME string   `json:"name"`
	TIP  string   `json:"type"`
}

func GetDistEtnii(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Extragere date din GET
	idScoala := c.Query("id_scoala")
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Verifiare daca userul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin pentru aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin pentru aceasta scoala"})
		return
	}
	etnii := map[string][]int{}
	clase := []string{}
	q := `select distinct etnie
		from elev
		where id_scoala = ?`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var etnie string
		if err := rows.Scan(&etnie); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			etnii[etnie] = []int{}
		}
	}
	if len(etnii) == 0 {
		c.IndentedJSON(http.StatusOK, false)
		return
	}
	q = `select id_clasa,etnie, count(*)
		from elev
		where id_scoala=?
		GROUP by id_clasa,etnie`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var etnie string
		var clasa string
		var cnt int
		if err := rows.Scan(&clasa, &etnie, &cnt); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			if len(clase) == 0 {
				clase = append(clase, clasa)
				etnii[etnie] = append(etnii[etnie], cnt)
				continue
			}
			if clase[len(clase)-1] == clasa {
				etnii[etnie] = append(etnii[etnie], cnt)
				continue
			}
			for key := range etnii {
				if len(etnii[key]) != len(clase) {
					etnii[key] = append(etnii[key], 0)
				}
			}
			clase = append(clase, clasa)
			etnii[etnie] = append(etnii[etnie], cnt)
		}
	}
	var date []Data = []Data{}
	for key := range etnii {
		date = append(date, Data{
			X:    clase,
			Y:    etnii[key],
			NUME: key,
			TIP:  "bar",
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"barmode": "stack",
	}})
}
