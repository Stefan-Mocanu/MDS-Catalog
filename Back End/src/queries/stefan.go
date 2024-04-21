package queries

import (
	"backend/queries/stefan"

	"github.com/gin-gonic/gin"
)

func Route_stefan(router gin.IRouter) {
	router.GET("/stefan", stefan.STEVE)
	router.GET("/exemplu", stefan.GetEXEMPLU)
	router.POST("/signin", stefan.Signin)
	router.POST("login", stefan.Login)
	router.POST("logout", stefan.Logout)
}
