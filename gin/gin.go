package gin

import (
	"context"
	"errors"
	"github.com/AndySu1021/go-util/log"
	"go.uber.org/fx"
	"golang.org/x/time/rate"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

var limiter *rate.Limiter

func NewGin(lc fx.Lifecycle, cfg *Config) *gin.Engine {
	gin.SetMode(cfg.Mode)

	limiter = initRateLimiter()

	var e = gin.New()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: e,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Logger.Infof("Starting gin server, listen on %s.", cfg.Port)
			var err error
			go func() {
				if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Logger.Errorf("Fail to run gin server, err: %+v", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			defer log.Logger.Info("Stopping gin HTTP server.")
			c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			return server.Shutdown(c)
		},
	})

	e.Use(RecoveryMiddleware())
	e.Use(LogRequest())
	e.Use(RequestIDMiddleware())
	e.Use(RateLimiterMiddleware())
	e.Use(CORS())

	pprof.Register(e)
	return e
}

func initRateLimiter() *rate.Limiter {
	return rate.NewLimiter(rate.Every(100*time.Microsecond), 10000)
}
