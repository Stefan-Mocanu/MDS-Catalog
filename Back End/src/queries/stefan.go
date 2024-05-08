package queries

import (
	"backend/queries/stefan"

	"github.com/gin-gonic/gin"
)

func Route_stefan(router gin.IRouter) {
	//router.GET("/stefan", stefan.STEVE)
	//router.POST("/exemplu", stefan.GetEXEMPLU)

	/*
		Cerere pentru crearea unu cont de utilizator
	*/
	router.POST("/signup", stefan.Signup)
	/*
		Cerere pentru autentificarea pe platforma
	*/
	router.POST("/login", stefan.Login)
	/*
		Cerere pentru terminarea unei sesiuni pe platforma
	*/
	router.POST("/logout", stefan.Logout)
	/*
		Cerere pentru verficarea faptului ca o sesiune este activa
	*/
	router.GET("/sessionActive", stefan.IsSessionActive)
	/*
		Cerere pentru obtinerea rolurilor unui utilizator
	*/
	router.GET("/getRoluri", stefan.GetRoluri)
	/*
		Inserarea materiilor si a incadrarilor in baza de date
	*/
	router.POST("/insertMaterie", stefan.Info_Materii)
	router.POST("/insertIncadrare", stefan.Info_incadrare)
	/*
		Obtinerea din baza de date a tokenurilor de linkuire cont-rol pentru profesor si elev/parinte
	*/
	router.GET("/csvProfesor", stefan.CreateCSVprofesor)
	router.GET("/csvElev", stefan.CreateCSVelev)
}
