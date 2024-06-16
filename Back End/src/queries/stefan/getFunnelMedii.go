package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetFunnelMedii(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din GET
	idScoala := c.Query("id_scoala")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin in aceasta scoala"})
		return
	}
	praguri := []int{5, 6, 7, 8, 9}
	valori := []int{}
	labels := []string{}
	for _, value := range praguri {
		q := `SELECT COUNT(*)
FROM (
    SELECT cls, elv, etn, AVG(medie) AS avg_medie
    FROM (
        SELECT 
            n.id_clasa AS cls, 
            n.id_elev AS elv, 
            e.etnie AS etn,
            AVG(n.nota) AS medie
        FROM 
            note n
        JOIN 
            elev e ON n.id_scoala = e.id_scoala 
                  AND n.id_clasa = e.id_clasa 
                  AND n.id_elev = e.id_elev
        WHERE 
            n.id_scoala = ?
        GROUP BY 
            n.id_clasa, n.id_elev, e.etnie, n.nume_disciplina
    ) AS inner_query
    GROUP BY 
        cls, elv, etn
) AS outer_query
WHERE avg_medie > ?;`
		var cnt int
		err := db.QueryRow(q, idScoala, value).Scan(&cnt)
		if err != nil {
			fmt.Println("Eroare: ", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
			return
		}
		valori = append(valori, cnt)
		labels = append(labels, "Numar de elevi<br>cu media peste "+fmt.Sprint(value))
		if cnt == 0 {
			break
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": []map[string]interface{}{{
		"type":      "funnel",
		"y":         labels,
		"x":         valori,
		"hoverinfo": "x+percent previous+percent initial",
	}}, "layout": map[string]interface{}{
		"margin": map[string]interface{}{"l": 150},
		"width":  600,
		"height": 500,
	}})
}
