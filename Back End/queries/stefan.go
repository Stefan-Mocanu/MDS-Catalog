package queries

import (
	"backend/queries/stefan"

	"github.com/gin-gonic/gin"
)

func Route_stefan(router gin.IRouter) {
	router.GET("/stefan", stefan.STEVE)
	router.GET("/exemplu", stefan.GetEXEMPLU)
	//router.GET/POST ....
}
