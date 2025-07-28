package main

import (
	"fmt"

	"github.com/Sudhir0302/secure_vault/services/storage/config"
	"github.com/Sudhir0302/secure_vault/services/storage/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("storage-service")
	config.Load()
	app := gin.New()
	routes.ConfigRoutes(app)
}
