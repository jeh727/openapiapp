package main

import (
	"log"
	"time"

	"github.com/jeh727/sampleapp/internal/app"
)

func main() {
	err := app.RunServer(":8080", 3*time.Second)
	if err != nil {
		log.Printf("server exiting with error: %v", err)
	}
}
