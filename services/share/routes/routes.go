package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(app *gin.Engine) {

	app.GET("/test", Test)
	app.POST("/addshare", AddShare)
	app.GET("/getshare", GetShare)

	app.Run(":8082")
}
