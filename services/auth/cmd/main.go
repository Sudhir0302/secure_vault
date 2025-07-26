package main

import (
	"fmt"

	"github.com/Sudhir0302/secure_vault.git/services/auth/config"
)

func main() {
	fmt.Println("auth-service")
	config.Load()
}
