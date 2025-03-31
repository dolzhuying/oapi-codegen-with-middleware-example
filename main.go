package main

import (
	"log"
	"oapi-codegen-with-middleware-example/router"
)

func main() {
	

	r := router.SetupRouter()
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
