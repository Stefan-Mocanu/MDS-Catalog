package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetSunBurstIncadrare(c *gin.Context) {
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
	q := `select nume
		from discipline
		where id_scoala=?`
	rows, err := db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	var ids = []string{}
	var parents = []string{}
	var labels = []string{}
	for rows.Next() {
		var nume string
		if err := rows.Scan(&nume); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			ids = append(ids, nume)
			parents = append(parents, "")
			labels = append(labels, strings.Replace(nume, " ", "<br>", -1))
		}
	}
	q = `SELECT distinct concat(p.nume," ",p.prenume),i.nume_disciplina
from profesor p, incadrare i
where p.id_scoala = ?
and p.id_scoala = i.id_scoala
and p.id = i.id_profesor`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var nume string
		var materie string
		if err := rows.Scan(&nume, &materie); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			ids = append(ids, materie+" - "+nume)
			parents = append(parents, materie)
			labels = append(labels, strings.Replace(nume, " ", "<br>", -1))
		}
	}
	q = `SELECT concat(p.nume," ",p.prenume),i.nume_disciplina,i.id_clasa
from profesor p, incadrare i
where p.id_scoala = ?
and p.id_scoala = i.id_scoala
and p.id = i.id_profesor`
	rows, err = db.Query(q, idScoala)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var nume string
		var materie string
		var clasa string
		if err := rows.Scan(&nume, &materie, &clasa); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			ids = append(ids, materie+" - "+nume+" - "+clasa)
			parents = append(parents, materie+" - "+nume)
			labels = append(labels, strings.Replace(clasa, " ", "<br>", -1))
		}
	}
	type data struct {
		TYPE    string   `json:"type"`
		IDS     []string `json:"ids"`
		PARENTS []string `json:"parents"`
		LABELS  []string `json:"labels"`
	}
	x := data{
		TYPE:    "sunburst",
		IDS:     ids,
		PARENTS: parents,
		LABELS:  labels,
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": x, "layout": map[string]interface{}{
		"margin": map[string]interface{}{"t": 0, "b": 0, "l": 0, "r": 0},
	}})
}
