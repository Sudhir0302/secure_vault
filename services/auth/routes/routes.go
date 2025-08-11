package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(app *gin.Engine) {

	//prefix each endpoint by /auth
	auth := app.Group("/auth/")

	auth.GET("/test", Test)
	auth.POST("/signup", Signup)
	auth.POST("/signin", Signin)
	app.Run(":8080")
}
