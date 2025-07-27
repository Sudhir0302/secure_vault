package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(app *gin.Engine) {

	app.GET("/test", Test)
	app.POST("/upload", UploadFile)

	app.Run(":8081")
}
