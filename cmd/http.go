package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhasnanr/e-wallet/bootstrap"
	"github.com/mhasnanr/e-wallet/internal/middleware"
)

func ServeHTTP() {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware(bootstrap.Log))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is healthy"})
	})

	server := &http.Server{Addr: ":" + bootstrap.GetEnv("HTTP_PORT", "8080"), Handler: r}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server stopped")
	}
}
