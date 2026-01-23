package main

import (
	"fmt"
	"log"

	"github.com/devoraq/Obfuscatorium_backend/internal/app"
	"github.com/devoraq/Obfuscatorium_backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Sprintf("Config error: %s", err)
	}
	
	app := app.New(cfg)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
