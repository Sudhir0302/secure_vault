package routes

import (
	"github.com/Sudhir0302/secure_vault.git/services/storage/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(app *gin.Engine) {

	app.GET("/test", Test)
	auth := app.Group("/api")

	auth.Use(middlewares.Verify())
	{
		auth.POST("/upload", UploadFile)
		auth.GET("/getfile", GetFile)
	}

	app.Run(":8081")
}
