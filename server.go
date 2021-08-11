package main

import (
	"apigolang/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file to set security or credentials information
	err := godotenv.Load()
	if err != nil {
		panic("Env file is required")
	}
	server := gin.Default()
	new(router.Router).Boot(server)
	server.Run()
}
