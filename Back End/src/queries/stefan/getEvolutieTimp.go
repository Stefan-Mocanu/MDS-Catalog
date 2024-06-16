package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetEvolutieTimp(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	//Verificare daca userul este logat
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din GET
	id_scoala := c.Query("id_scoala")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Parinte",
		SCOALA: id_scoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este admin in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este admin in aceasta scoala"})
		return
	}
	q := `SELECT date_format(n.data, '%Y-%M')           AS de,
		AVG(n.nota)                            AS n_elev,
		(SELECT AVG(nn.nota) AS average_nota
			FROM note nn
			WHERE id_scoala = ?
			AND id_clasa = ?
			AND nume_disciplina = ?
			AND YEAR(nn.data) = YEAR(n.data)
			AND MONTH(nn.data) = MONTH(n.data)) AS overall
	FROM note n
	WHERE n.id_scoala = ?
	AND n.id_clasa = ?
	AND n.id_elev = ?
	AND n.nume_disciplina = ?
	GROUP BY YEAR(n.data),
			MONTH(n.data);`

	id_clasa := c.Query("id_clasa")
	nume_disciplina := c.Query("nume_disciplina")
	id_elev := c.Query("id_elev")

	rows, err := db.Query(q, id_scoala, id_clasa, nume_disciplina, id_scoala, id_clasa, id_elev, nume_disciplina)

	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()

	a_e := []float64{}
	a_m := []float64{}
	d_t := []string{}

	for rows.Next() {
		var n_e, n_m float64
		var data string
		if err := rows.Scan(&data, &n_e, &n_m); err != nil {
			fmt.Println("Error: ", err)
		} else {
			a_e = append(a_e, n_e)
			a_m = append(a_m, n_m)
			d_t = append(d_t, data)
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H(map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"x":    d_t,
				"y":    a_e,
				"type": "scatter",
				"name": "Medie elev",
			},
			map[string]interface{}{
				"x":    d_t,
				"y":    a_m,
				"type": "scatter",
				"name": "Medie clasă",
			},
		},
		"layout": map[string]interface{}{},
	}))
}
