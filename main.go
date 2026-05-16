package main

import (
	"context"
	"log"

	"github.com/mhasnanr/ewallet-transaction/bootstrap"
	"github.com/mhasnanr/ewallet-transaction/cmd"
)

func main() {
	if err := bootstrap.SetupZapLogger(); err != nil {
		log.Fatal("failed to initialize logger")
	}

	if err := bootstrap.SetupConfig(".env"); err != nil {
		log.Fatalf("failed to load config file")
	}

	ctx := context.Background()
	serviceName := bootstrap.GetEnv("SERVICE_NAME", "ewallet-transaction")
	endpoint := bootstrap.GetEnv("OTELExporterOTLPEndpoint", "http://otel-collector:4318")
	shutdown, err := bootstrap.SetupOTel(ctx, serviceName, endpoint)
	if err != nil {
		log.Fatalf("failed to setup otel: %v", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown otel: %v", err)
		}
	}()

	bootstrap.SetupDatabase()

	go cmd.ServeGRPC()
	cmd.ServeHTTP(bootstrap.DB)
}
