package main

import (
	"fmt"

	"github.com/Sudhir0302/secure_vault/services/auth/config"
	"github.com/Sudhir0302/secure_vault/services/auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("auth-service")
	config.Load()

	// gin.SetMode(gin.ReleaseMode)  //disables default logs

	app := gin.New()
	routes.ConfigRoutes(app)
}
