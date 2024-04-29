package main

import (
	"fmt"
	"os"
	"rate-limiter/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("Fixed port to 5000")
	}

	fmt.Println("Listening port: " + port)
	server.New().Run(":" + port)
}
