package main

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	return router
}
func main() {
	router := SetUpRouter()

	SetupRoutes(router)
	_ = router.Run(":8080")
}
