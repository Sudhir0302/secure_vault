package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(app *gin.Engine) {

	wrap := app.Group("/share/")
	wrap.GET("/test", Test)

	// auth := app.Group("/api")
	// auth.Use(middlewares.Verify())
	// {
	// 	auth.POST("/addshare", AddShare)
	// 	auth.GET("/getshare", GetShare)
	// }
	wrap.POST("/addshare", AddShare)
	wrap.GET("/getshare", GetShare)

	app.Run(":8082")
}
