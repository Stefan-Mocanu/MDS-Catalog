package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetEvolutie(c *gin.Context) {
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
	idClasa := c.Query("id_clasa")
	materie := c.Query("materie")
	data := c.Query("data") //yyyy-mm-dd
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}
	q := `select CONCAT(e.nume," ",e.prenume), 
                ifnull((select AVG(nota)
                from note n2
                where id_scoala = ?
                and id_clasa = ?
                and n2.id_elev = n.id_elev
                and nume_disciplina=?) - (select AVG(nota)
                from note n1
                where id_scoala = ?
                and id_clasa = ?
                and n1.id_elev = n.id_elev
                and data<?
                and nume_disciplina=?),0) 
		from note n,elev e
		where n.id_scoala = ?
		and n.id_clasa = ?
		and n.id_clasa = e.id_clasa
		and n.id_scoala = e.id_scoala
		and n.id_elev = e.id_elev
		GROUP by 1`
	rows, err := db.Query(q, idScoala, idClasa, materie, idScoala, idClasa, data, materie, idScoala, idClasa)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	numeElevi := []string{}
	diferente := []float64{}

	defer rows.Close()
	for rows.Next() {
		var nume string
		var diferenta float64
		if err := rows.Scan(&nume, &diferenta); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			numeElevi = append(numeElevi, nume)
			diferente = append(diferente, diferenta)
		}
	}
	date := []Data2{{
		X:   numeElevi,
		Y:   diferente,
		TIP: "bar",
	}}
	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"Title": "Diferenta mediilor inainte de " + data + " si dupa",
	}})
}

func GetBoxNoteActivitate(c *gin.Context) {
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
	idClasa := c.Query("id_clasa")
	materie := c.Query("materie")
	//Verificare daca useul este ADMIN
	if (!VerificareRol(Rol{
		ROL:    "Administrator",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}
	q := `select nota, EXTRACT(YEAR FROM data) AS year, EXTRACT(MONTH FROM data) AS month
		from note
		where id_scoala = ?
		and id_clasa = ?
		and nume_disciplina = ?
		order by data`
	rows, err := db.Query(q, idScoala, idClasa, materie)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	note := []int{}
	date_note := []string{}
	defer rows.Close()
	for rows.Next() {
		var nota int
		var an, luna string
		if err := rows.Scan(&nota, &an, &luna); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			note = append(note, nota)
			date_note = append(date_note, an+"/"+luna)
		}
	}
	q = `select valoare, EXTRACT(YEAR FROM data) AS year, EXTRACT(MONTH FROM data) AS month
		from activitate
		where id_scoala = ?
		and id_clasa = ?
		and nume_disciplina = ?
		and not valoare = 0
		order by data`
	rows, err = db.Query(q, idScoala, idClasa, materie)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": err})
		return
	}
	activitate := []int{}
	date_activ := []string{}
	defer rows.Close()
	for rows.Next() {
		var nota int
		var an, luna string
		if err := rows.Scan(&nota, &an, &luna); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			activitate = append(activitate, nota)
			date_activ = append(date_activ, an+"/"+luna)
		}
	}
	type data struct {
		X    []string `json:"x"`
		Y    []int    `json:"y"`
		NAME string   `json:"name"`
		TIP  string   `json:"type"`
	}
	date := []data{}
	date = append(date, data{
		X:    date_note,
		Y:    note,
		NAME: "Note",
		TIP:  "box",
	})
	date = append(date, data{
		X:    date_activ,
		Y:    activitate,
		NAME: "Activitate",
		TIP:  "box",
	})
	c.IndentedJSON(http.StatusOK, gin.H{"data": date, "layout": map[string]interface{}{
		"yaxis": map[string]interface{}{
			"title":    "Distributia notelor si activitatii pe luni",
			"zeroline": false,
		},
		"boxmode": "group",
	}})
}
