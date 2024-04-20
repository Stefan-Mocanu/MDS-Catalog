package main

import (
	"backend/queries"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/exemplu", queries.GetEXEMPLU)
	router.Run("localhost:9000")
}
