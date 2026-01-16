package main

import (
	"log"

	"github.com/devoraq/Obfuscatorium_backend/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
