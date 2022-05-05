package db

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type DatabaseType string

func (d DatabaseType) String() string {
	return string(d)
}

const (
	MySQL  DatabaseType = "mysql"
	SQLite DatabaseType = "sqlite"
)

type Config struct {
	Debug              bool         `mapstructure:"debug"`
	Type               DatabaseType `mapstructure:"type"`
	Host               string       `mapstructure:"host"`
	Port               int          `mapstructure:"port"`
	Username           string       `mapstructure:"username"`
	Password           string       `mapstructure:"password"`
	DBName             string       `mapstructure:"db_name"`
	MaxIdleConnections int          `mapstructure:"max_idle_connections"`
	MaxOpenConnections int          `mapstructure:"max_open_connections"`
	MaxLifetimeSec     int          `mapstructure:"max_lifetime_sec"`
	WithColor          bool         `mapstructure:"with_color"`
}

func GetConnectionStr(cfg *Config) (connectionString string, err error) {
	switch cfg.Type {
	case MySQL:
		connectionString = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&multiStatements=true&parseTime=true", cfg.Username, cfg.Password, cfg.Host+":"+strconv.Itoa(cfg.Port), cfg.DBName)
	case SQLite:
		if cfg.Host == "" {
			connectionString = path.Join(os.Getenv("PROJ_DIR"), "test/.data", "sqlite.db?cache=shared")
		} else {
			connectionString = cfg.Host
		}
	default:
		return "", errors.New("not support driver")
	}

	return
}

func NewDatabase(cfg *Config) (db *sql.DB, err error) {
	dsn, err := GetConnectionStr(cfg)
	if err != nil {
		return nil, err
	}

	db, err = sql.Open(cfg.Type.String(), dsn)
	if err != nil {
		log.Error().Msgf("fail to open connection, err: %+v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("database ping success")

	if cfg.MaxIdleConnections != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
	} else {
		db.SetMaxIdleConns(2)
	}

	if cfg.MaxOpenConnections != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConnections)
	} else {
		db.SetMaxOpenConns(5)
	}

	if cfg.MaxLifetimeSec != 0 {
		db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetimeSec) * time.Second)
	} else {
		db.SetConnMaxLifetime(1 * time.Hour)
	}

	return db, err
}