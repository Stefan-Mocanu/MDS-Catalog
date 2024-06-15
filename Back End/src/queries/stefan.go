package queries

import (
	"backend/queries/stefan"

	"github.com/gin-gonic/gin"
)

func Route_stefan(router gin.IRouter) {
	router.GET("/stefan", stefan.STEVE)
	router.GET("/exemplu", stefan.GetEXEMPLU)
	router.POST("/signup", stefan.Signup)
	/*
		Cerere pentru autentificarea pe platforma
		Primeste parametrii din POST: email, parola
	*/
	router.POST("/login", stefan.Login)
	/*
		Cerere pentru terminarea unei sesiuni pe platforma
		Nu necesita parametrii
	*/
	router.POST("/logout", stefan.Logout)
	/*
		Cerere pentru verficarea faptului ca o sesiune este activa
		Nu necesita parametrii
	*/
	router.GET("/sessionActive", stefan.IsSessionActive)
	/*
		Cerere pentru obtinerea rolurilor unui utilizator
		Nu necesita parametrii
	*/
	router.GET("/getRoluri", stefan.GetRoluri)
	/*
		Inserarea materiilor si a incadrarilor in baza de date
		Primesc ca parametrii din POST: id_scoala, csv_file(fisier .CSV)
	*/
	router.POST("/insertMaterie", stefan.Info_Materii)
	router.POST("/insertIncadrare", stefan.Info_incadrare)
	/*
		Obtinerea din baza de date a tokenurilor de linkuire cont-rol pentru profesor si elev/parinte
		Primeste ca parametrii din GET: id_scoala
	*/
	router.GET("/csvProfesor", stefan.CreateCSVprofesor)
	router.GET("/csvElev", stefan.CreateCSVelev)
	/*
		Utilizarea unui token pentru linkuirea unui cont cu un profesor/elev/parinte
		Primeste ca parametrii din POST: rol("elev"/"profesor"/"parinte"), token
	*/
	router.POST("/alaturare", stefan.Alaturare)
	/*
		Obtinerea catalogului din postura de elev
		Primeste ca parametrii din GET: id_scoala, id_clasa
	*/
	router.GET("/viewCatalogElev", stefan.View_note_elev)
	/*
		Obtinerea catalogului din postura de parinte
		Primeste ca parametrii din GET: id_scoala, id_clasa, id_elev
	*/
	router.GET("/viewCatalogParinte", stefan.View_note_parinte)
	/*
		Obtinerea copiilor inregistrati la o scoala ai unui parinte
		Primeste ca parametru din GET: id_scoala
	*/
	router.GET("/getElevi", stefan.GetElevi)
	/*
		Adaugare admin pentru o scoala cu un cont de admin
		Primeste ca parametrii in POST: id_scoala, id_cont
	*/
	router.POST("/adaugaAdmin", stefan.AdaugaAdmin)
	/*
		Obtinerea claselor la care este inregistrat un elev intr-o scoala
		Primeste ca parametrii din GET: id_scoala
	*/
	router.GET("/getClase", stefan.GetClasa)
	/*
		Obtinerea parametrilor de realizare a unui grafic barplot cu repartitia de etnii din fiecare clasa
		Primeste ca parametru din GET: id_scoala
	*/
	router.GET("/getDistEtnii", stefan.GetDistEtnii)
	/*
		Obtinerea parametrilor de realizare a unui grafic barplot cu deviatia standard a mediilor elevilor
		Primeste ca parametru din GET: id_scoala
	*/
	router.GET("/getDevStdNote", stefan.GetDevStdNote)
	/*
		Trimiterea unui feedback de catre elev pentru profesor
		Primeste ca parametrii din POST: id_scoala, id_clasa, materie, tip(0/1), continut(string)
	*/
	router.POST("/addFeedbackElevProf", stefan.FeedbackElevProfesor)
	/*
		Trimiterea unui feedback de catre profesor pentru elev
		Primeste ca parametrii din POST: id_scoala, id_clasa, materie, tip(0/1), continut(string), id_elev
	*/
	router.POST("/addFeedbackProfElev", stefan.FeedbackProfesorElev)
	/*
		Obtinerea feedbackului primit de un profesor de la elevii unei clase la care preda
		Primeste ca parametrii din GET: id_scoala, id_clasa, materie
	*/
	router.GET("/getFeedbackProf", stefan.GetFeedback)
	/*
		Obtinerea materiilor dintr-o clasa la care un elev face parte
		Primeste ca parametrii din GET: id_scoala, id_clasa
	*/
	router.GET("/getMaterii", stefan.GetMaterii)
	/*
		Obtinerea parametrilor de realizare a unui grafic barplot pentru comparatia mediilor pana la o anumita data si dupa aceasta
		Primeste ca parametrii din GET: id_scoala, id_clasa, materie, data
		Acest grafic este pentru profesor
	*/
	router.GET("/getEvolutie", stefan.GetEvolutie)
	/*
		Obtinerea parametrilor de realizare a unui grafic boxplot pentru vizualizarea notelor si activitatii
		Primeste ca parametrii din GET: id_scoala, id_clasa, materie
		Acest grafic este pentru profesor
	*/
	router.GET("/getNoteActivitate", stefan.GetBoxNoteActivitate)

	router.GET("/GetMedieClasament", stefan.GetMedieClasament)
}
