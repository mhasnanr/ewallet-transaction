package main

import (
	"log"

	"github.com/mhasnanr/e-wallet/bootstrap"
)

func main() {
	if err := bootstrap.SetupZapLogger(); err != nil {
		log.Fatal("failed to initialize logger")
	}

	if err := bootstrap.SetupConfig(".env"); err != nil {
		log.Fatalf("failed to load config file")
	}
}
