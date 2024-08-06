package router

import (
	"api_crud_blog/app/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter is a function to set router
func SetupRouter(mode string) *gin.Engine {
	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Connect to database and initial migration
	database.ConnectDB()
	database.InitialMigration()

	router.Use(cors.Default())

	return router
}