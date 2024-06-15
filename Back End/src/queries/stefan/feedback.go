package stefan

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FeedbackElevProfesor(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.PostForm("id_scoala")
	idClasa := c.PostForm("id_clasa")
	materie := c.PostForm("materie")
	tip := c.PostForm("tip")
	content := c.PostForm("continut")
	if (!VerificareRol(Rol{
		ROL:    "Elev",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este elev in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este elev in aceasta scoala"})
		return
	}
	q := `select id_elev
		from elev
		where id_scoala = ?
		and id_clasa = ?
		and id_cont_elev = ?`
	idElev := 0
	err := db.QueryRow(q, idScoala, idClasa, ver).Scan(&idElev)
	switch {

	case err == sql.ErrNoRows:
		fmt.Printf("Eroare: %v", err)

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Nu exista acest elev."})
	case err != nil:
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
	}
	q = `insert into feedback(id_scoala,nume_disciplina,id_clasa,id_elev,content,data,tip,directie)
		values(?,?,?,?,?,SYSDATE(),?,?)`
	_, err = db.Exec(q, idScoala, materie, idClasa, idElev, content, tip, 1)
	if err != nil {
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
	}
}

func FeedbackProfesorElev(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.PostForm("id_scoala")
	idClasa := c.PostForm("id_clasa")
	materie := c.PostForm("materie")
	tip := c.PostForm("tip")
	content := c.PostForm("continut")
	idElev := c.PostForm("id_elev")
	if (!VerificareRol(Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}
	q := `insert into feedback(id_scoala,nume_disciplina,id_clasa,id_elev,content,data,tip,directie)
		values(?,?,?,?,?,SYSDATE(),?,?)`
	_, err := db.Exec(q, idScoala, materie, idClasa, idElev, content, tip, 0)
	if err != nil {
		fmt.Printf("Eroare: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
	}
}

func GetFeedback(c *gin.Context) {
	var db *sql.DB = database.InitDb()
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	idScoala := c.Query("id_scoala")
	idClasa := c.Query("id_clasa")
	materie := c.Query("materie")
	if (!VerificareRol(Rol{
		ROL:    "Profesor",
		SCOALA: idScoala,
		ID:     ver,
	})) {
		fmt.Println("Userul nu este profesor in aceasta scoala")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este profesor in aceasta scoala"})
		return
	}

	var feedback = []Feedback{}
	q := `select content, data, tip
		from feedback
		where id_scoala = ?
		and id_clasa = ?
		and directie = 1
		and nume_disciplina = ?
		order by data`
	rows, err := db.Query(q, idScoala, idClasa, materie)
	if err != nil {
		fmt.Println("Eroare: ", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Eroare": "Alta eroare"})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var data, content string
		var tip bool
		if err := rows.Scan(&content, &data, &tip); err != nil {
			fmt.Println("Eroare: ", err)
		} else {
			feedback = append(feedback, Feedback{
				CONTENT: content,
				DATA:    data,
				TIP:     tip,
			})
		}
	}
	c.IndentedJSON(http.StatusOK, feedback)
}
