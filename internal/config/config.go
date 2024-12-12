package config

import (
	"github.com/Azaliya1995/music_library/pkg/log"
	"github.com/Azaliya1995/music_library/version"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	ServerConfig   ServerConfig   `envPrefix:"SERVER_"`
	LogConfig      log.Config     `envPrefix:"LOG_"`
	DatabaseConfig DatabaseConfig `envPrefix:"DB_"`
}

type ServerConfig struct {
	Listen            string        `env:"LISTEN" envDefault:":8080"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" envDefault:"5s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
	MaxHeaderSize     int           `env:"MAX_HEADER_SIZE" envDefault:"5000"`
	GracefulShutdown  time.Duration `env:"GRACEFUL_SHUTDOWN" envDefault:"5s"`
	DebugGin          bool          `env:"DEBUG" envDefault:"false"`
	RemoveExtraSlash  bool          `env:"REMOVE_EXTRA_SLASH" envDefault:"false"`
}

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	cfg.LogConfig.Fields.Version = version.Version
	cfg.LogConfig.Fields.CommitHash = version.CommitHash
	cfg.LogConfig.Fields.CommitTime = version.CommitTime

	return &cfg, nil
}
