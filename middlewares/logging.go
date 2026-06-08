package middlewares

import (
	"log"
	"time"
	"github.com/beego/beego/v2/server/web/context"
)

type FilterContext = context.Context // Helper alias reference

func RequestLogger() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		startTime := time.Now()
		
	// Add a custom header to all responses for demonstration

		ctx.ResponseWriter.Header().Set("X-Server-Framework", "Beego")
		
		// Defer a function to log the request details and duration after the request is processed
		defer func() {
			duration := time.Since(startTime)
			log.Printf("[METRIC] Method: %s | URL: %s | Duration: %s", 
				ctx.Input.Method(), 
				ctx.Input.URL(), 
				duration,
			)
		}()
	}
}