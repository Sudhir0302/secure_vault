package routes

import (
	"github.com/Sudhir0302/secure_vault/services/share/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(app *gin.Engine) {

	app.GET("/test", Test)

	auth := app.Group("/api")
	auth.Use(middlewares.Verify())
	{
		auth.POST("/addshare", AddShare)
		auth.GET("/getshare", GetShare)
	}

	app.Run(":8082")
}
