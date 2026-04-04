package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "server is healthy"})
	})

	server := &http.Server{Addr: ":8080", Handler: r}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server stopped")
	}
}
