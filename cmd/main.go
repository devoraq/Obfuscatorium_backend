package main

import (
	"log"

	"github.com/user/obfuscatorium/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
