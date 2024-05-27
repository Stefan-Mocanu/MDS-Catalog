package main

import (
	"backend/queries"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	queries.Route_gabi(router)
	queries.Route_stefan(router)
	// queries.Route_gabi(router)
	router.Run("localhost:9000")

}
