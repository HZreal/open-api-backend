package main

import (
	"open-api-backend/routes"
	"open-api-backend/rpc/server"
)

func main() {
	//
	server.StartGPRC()

	//
	routes.StartGinServer()
}
