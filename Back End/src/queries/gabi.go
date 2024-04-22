package queries

import (
	"backend/queries/gabi"

	"github.com/gin-gonic/gin"
)

func Route_gabi(router gin.IRouter) {
	router.GET("/gabi", gabi.GABI)
	router.GET("/exemplu1", gabi.GetEXEMPLU1)
}
