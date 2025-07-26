package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(app *gin.Engine) {
	app.GET("/test", Test)
	app.POST("/signup", Signup)
	app.Run(":8080")
}
