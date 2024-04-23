package queries

import (
	"backend/queries/gabi"

	"github.com/gin-gonic/gin"
)

func Route_gabi(router gin.IRouter) {
	router.GET("/gabi", gabi.GABI)
	router.GET("/exemplu1", gabi.GetEXEMPLU1)
	router.POST("/inserareScoala", gabi.Inregistrare_Instit_Invat)
	router.POST("/inserareClasa", gabi.Info_clase)
	router.POST("/inserareProfesor", gabi.Info_profesor)
	router.POST("/inserareElev", gabi.Info_elev)

}
