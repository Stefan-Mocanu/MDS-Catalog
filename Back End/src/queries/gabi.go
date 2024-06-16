package queries

import (
	"backend/queries/gabi"

	"github.com/gin-gonic/gin"
)

func Route_gabi(router gin.IRouter) {
	// Rute pentru diverse funcționalități ale modulului "gabi"

	// Răspuns simplu pentru ruta de bază "/gabi"
	router.GET("/gabi", gabi.GABI)

	// Răspuns pentru ruta "/exemplu1"
	router.GET("/exemplu1", gabi.GetEXEMPLU1)

	// Inserarea unei noi școli în baza de date
	// Primesc ca parametrii din POST: nume_scoala, adresa, telefon
	router.POST("/inserareScoala", gabi.Inregistrare_Instit_Invat)

	// Inserarea unei noi clase în baza de date
	// Primesc ca parametrii din POST: id_scoala, csv_file (fișier .CSV)
	router.POST("/inserareClasa", gabi.Info_clase)

	// Inserarea unui nou profesor în baza de date
	// Primesc ca parametrii din POST: id_scoala, nume, prenume, materie, email
	router.POST("/inserareProfesor", gabi.Info_profesor)

	// Inserarea unui nou elev în baza de date
	// Primesc ca parametrii din POST: id_scoala, nume, prenume, clasa, email
	router.POST("/inserareElev", gabi.Info_elev)

	// Oferirea feedbackului unui elev de catre un profesor"
	// Primesc ca parametrii din POST : ID-ul profesorului, ID-ul școlii.
	// Apoi se primesc ca parametrii: nume_disciplina, id_clasa, id_elev
	router.POST("/feedback", gabi.Feedback)

	// Adăugarea unei note pentru un elev
	// Primesc ca parametrii din POST: nota, nume_disciplina, id_clasa, id_elev
	router.POST("/adaugareNota", gabi.Note)

	// Adăugarea unei prezențe pentru un elev
	// Primesc ca parametrii din POST: valoarea_prezentei(0 sau 1), id_profesor, id_scoala, nume_disciplina, id_clasa, id_elev
	router.POST("/adaugarePrezenta", gabi.Prezenta)

	// Obținerea claselor asociate unui profesor
	// Primesc ca parametrii din POST: id_profesor
	router.POST("/claseProfesor", gabi.Clase_Profesor)

	// Obținerea informațiilor despre elevii dintr-o clasă
	// Primesc ca parametrii din POST: id_clasa
	router.POST("/eleviClasa", gabi.EleviClasa)

	// Obținerea informațiilor despre elevii unui profesor de la o școală
	// Primesc ca parametrii din POST: id_profesor, id_scoala
	router.POST("/eleviProfesor", gabi.EleviProfesor)

	// Obținerea distribuției mediilor pe etnii
	// Răspunde la GET cu distribuția mediilor pe etnii pentru școala specificată
	router.GET("/getRepMediiEtnii", gabi.GetDistEtnii)

	// Obținerea distribuției mediilor pe etnii
	// Răspunde la GET cu distribuția mediilor pe genuri pentru școala specificată
	router.GET("/getRepMediiGen", gabi.GetDistGenuri)

	router.GET("/GetEvolNoteElevi", gabi.GetEvolutieNoteElevi)

	router.GET("/MediiClase", gabi.MediiClase)

	router.GET("/HM_MediiLuniA", gabi.HeatMapMediiLunileAnului)

	router.GET("/EleviPromov", gabi.EleviPromov)

	router.GET("/FeedbackuriProfesori", gabi.FeedbackuriProfesori)

	router.GET("/EthnicityClassStats", gabi.EthnicityClassStats)

	// router.GET("/FeedbackuriProfesori", gabi.FeedbackuriProfesori)

	router.GET("/getProfessorsFeedback2", gabi.ProfessorsFeedback2)

	router.GET("/getProfessorsFeedback", gabi.ProfessorsFeedback)

	router.GET("EthnicitySankey", gabi.EthnicitySankey)

	router.GET("/getEthnicityClassStats", gabi.EthnicityClassStats)

	router.GET("/ChatFeedbackProf", gabi.FeedbackuriProfesori)
}
