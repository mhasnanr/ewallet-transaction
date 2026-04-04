package main

import (
	"log"

	"github.com/mhasnanr/e-wallet/bootstrap"
	"github.com/mhasnanr/e-wallet/cmd"
)

func main() {
	if err := bootstrap.SetupZapLogger(); err != nil {
		log.Fatal("failed to initialize logger")
	}

	if err := bootstrap.SetupConfig(".env"); err != nil {
		log.Fatalf("failed to load config file")
	}

	dsn := bootstrap.GetEnv("CONNECTION_STRING", "")
	_, err := bootstrap.SetupDatabase(dsn)
	if err != nil {
		log.Fatalf("failed to connect database")
	}

	go cmd.ServeGRPC()
	cmd.ServeHTTP()
}
