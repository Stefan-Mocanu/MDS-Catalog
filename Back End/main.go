package main

import (
	"backend/queries"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	queries.Route_stefan(router)
	router.Run("localhost:9000")
}
