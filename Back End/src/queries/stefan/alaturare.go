package stefan

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Alaturare(c *gin.Context) {
	//Verificare daca sesiunea este activa
	ver := IsSessionActiveIntern(c)
	if ver < 0 {
		fmt.Println("Userul nu este logat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Userul nu este logat"})
		return
	}
	//Obtinere date din POST
	rol := c.PostForm("rol")
	if rol == "elev" {
		AlaturareElev(c)
		return
	}
	if rol == "parinte" {
		AlaturareParinte(c)
		return
	}
	if rol == "profesor" {
		AlaturareProf(c)
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "nu exista rolul"})

}
