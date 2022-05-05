package gin

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"net/http"
	"runtime"
	"strings"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				trace := make([]byte, 4096)
				runtime.Stack(trace, true)
				var msg string
				msg += fmt.Sprintf("%s\n", r)
				for i := 0; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					msg += fmt.Sprintf("%s:%d\n", file, line)
				}
				log.Error().Msgf("Unknown error: %s", msg)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    -1,
					"message": "unknown error",
					"data":    nil,
				})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !strings.Contains(c.Request.URL.Path, "health") {
			log.Ctx(c.Request.Context()).Info().
				Str("url", c.Request.URL.String()).
				Str("method", c.Request.Method).
				Interface("header", c.Request.Header).
				Bool("access_log", true).
				Msgf("access log")
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = xid.New().String()
		}

		logger := log.With().Str("trace_id", traceID).Logger()

		deviceID := c.GetHeader("X-Device-ID")

		ctx := context.WithValue(c.Request.Context(), "X-Trace-ID", traceID)
		ctx = context.WithValue(ctx, "X-Device-ID", deviceID)
		ctx = logger.WithContext(ctx)

		c.Request = c.Request.WithContext(ctx)
		c.Request.Header.Set("X-Trace-ID", traceID)
		c.Writer.Header().Set("X-Trace-ID", traceID)
		c.Next()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter == nil {
			c.Next()
		}

		err := limiter.Wait(c.Request.Context())
		if err != nil {
			return
		}

		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	return cors.New(corsConfig)
}
