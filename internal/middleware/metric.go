package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	httpRequestTotal metric.Int64Counter
)

func init() {
	meter := otel.Meter("transaction-service")
	var err error
	httpRequestTotal, err = meter.Int64Counter(
		"http_response_total",
		metric.WithDescription("Total number of HTTP responses by status code and handler."),
	)
	if err != nil {
		panic(err)
	}
}

func MetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		handler := c.FullPath()
		if handler == "" {
			handler = "unknown"
		}
		status := c.Writer.Status()

		httpRequestTotal.Add(c.Request.Context(), 1,
			metric.WithAttributes(
				attribute.String("handler", handler),
				attribute.String("code", strconv.Itoa(status)),
			),
		)
	}
}
