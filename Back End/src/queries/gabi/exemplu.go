package gabi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func GetEXEMPLU1(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "HELLO!")
}
