package queries

import (
	"backend/queries/gabi"

	"github.com/gin-gonic/gin"
)

func Route_gabi(router gin.IRouter) {
	router.GET("/gabi", gabi.GABI)
	router.GET("/exemplu1", gabi.GetEXEMPLU1)
	/*
		Inserarea unei noi școli în baza de date
		Primesc ca parametrii din POST: nume_scoala, adresa, telefon
	*/
	router.POST("/inserareScoala", gabi.Inregistrare_Instit_Invat)
	/*
		Inserarea unei noi clase în baza de date
		Primesc ca parametrii din POST: id_scoala, csv_file (fișier .CSV)
	*/
	router.POST("/inserareClasa", gabi.Info_clase)
	/*
		Inserarea unui nou profesor în baza de date
		Primesc ca parametrii din POST: id_scoala, nume, prenume, materie, email
	*/
	router.POST("/inserareProfesor", gabi.Info_profesor)
	/*
		Inserarea unui nou elev în baza de date
		Primesc ca parametrii din POST: id_scoala, nume, prenume, clasa, email
	*/
	router.POST("/inserareElev", gabi.Info_elev)

	router.POST("/feedback", gabi.Feedback)
}
