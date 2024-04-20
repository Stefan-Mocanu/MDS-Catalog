package stefan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Exemplu de functie HTTP
// FUNCTIILE INCEP CU LITERA MARE
func STEVE(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "STEFAN")
}
