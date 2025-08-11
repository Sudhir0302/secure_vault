package main

import (
	"fmt"
	"net/http"

	"github.com/Sudhir0302/secure_vault/services/api-gateway/gateway"
)

func main() {
	fmt.Println("api-gateway")

	//auth-gateway
	auth_service := gateway.ConfigGateway("http://localhost:8080")
	gateway.AuthRoute(auth_service)

	//share-gateway
	share_service := gateway.ConfigGateway("http://localhost:8082")
	gateway.ShareRoute(share_service)

	//store-gateway
	store_service := gateway.ConfigGateway("http://localhost:8081")
	gateway.StoreRoute(store_service)

	http.ListenAndServe(":8090", nil)
}
