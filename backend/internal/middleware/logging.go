package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Log request
		fmt.Printf("[REQUEST] %s %s\n", method, path)

		// Log request body for POST/PUT
		if method == "POST" || method == "PUT" {
			if c.Request.Body != nil {
				body, err := io.ReadAll(c.Request.Body)
				if err == nil && len(body) > 0 {
					fmt.Printf("[BODY] %s\n", string(body))
					// Restore body
					c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				}
			}
		}

		c.Next()

		// Log response
		duration := time.Since(start)
		fmt.Printf("[RESPONSE] %s %s - %d (%v)\n", method, path, c.Writer.Status(), duration)

		// Log errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				fmt.Printf("[ERROR] %s\n", e.Error())
			}
		}
	}
}