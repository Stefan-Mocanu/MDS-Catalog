package queries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func GetEXEMPLU(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "HELLO!")
}
