package zlog

import "github.com/rs/zerolog"

// Config ...
type Config struct {
	AppName     string        `mapstructure:"app_name"`
	Environment string        `mapstructure:"environment"` // local, dev, sitable, prod
	Level       zerolog.Level `mapstructure:"level"`       // Debug: 0, info: 1 , warn: 2, error: 3, fatal: 4, panic: 5
	Local       bool          `mapstructure:"local"`
}
