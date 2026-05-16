package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhasnanr/ewallet-transaction/bootstrap"
	"github.com/mhasnanr/ewallet-transaction/external"
	handler "github.com/mhasnanr/ewallet-transaction/internal/handler/http"
	"github.com/mhasnanr/ewallet-transaction/internal/middleware"
	"github.com/mhasnanr/ewallet-transaction/internal/repository"
	"github.com/mhasnanr/ewallet-transaction/internal/services"
	"gorm.io/gorm"
)

func ServeHTTP(db *gorm.DB) {
	r := gin.New()

	r.Use(
		middleware.LoggerMiddleware(bootstrap.Log),
		middleware.MetricMiddleware(),
		middleware.ErrorMiddleware(),
	)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is healthy"})
	})

	userGRPCClient, userConn, err := external.NewUserGRPC()
	if err != nil {
		log.Fatalf("failed to initialized user gRPC: %v", err)
	}
	defer userConn.Close()

	walletGRPCClient, walletConn, err := external.NewWalletGRPC()
	if err != nil {
		log.Fatalf("failed to initialized wallet gRPC: %v", err)
	}
	defer walletConn.Close()

	authMiddleware := middleware.NewAuthMiddleware(userGRPCClient)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, walletGRPCClient)
	transactionHandler := handler.NewTransactionHandler(transactionService, authMiddleware)

	healthCheckHandler := handler.NewHealthCheck()

	healthCheckHandler.RegisterRoute(r)
	transactionHandler.RegisterRoute(r)

	httpPort := bootstrap.GetEnv("HTTP_PORT", "8080")
	server := &http.Server{Addr: ":" + httpPort, Handler: r}

	fmt.Printf("http server is running on port %s...\n", httpPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server stopped")
	}
}
