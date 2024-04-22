package queries

import (
	"backend/queries/stefan"

	"github.com/gin-gonic/gin"
)

func Route_stefan(router gin.IRouter) {
	router.GET("/stefan", stefan.STEVE)
	router.GET("/exemplu", stefan.GetEXEMPLU)
	router.POST("/signup", stefan.Signup)
	router.POST("/login", stefan.Login)
	router.POST("/logout", stefan.Logout)
	router.GET("/sessionActive", stefan.IsSessionActive)
}
