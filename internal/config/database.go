package config

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Name            string        `env:"NAME"`
	Schema          string        `env:"SCHEMA"`
	Host            string        `env:"HOST"`
	User            string        `env:"USER"`
	Password        string        `env:"PASSWORD"`
	Port            int           `env:"PORT"`
	SSLMode         string        `env:"SSL_MODE"`
	ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `env:"CONN_MAX_IDLE_TIME"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS"`
}

func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
		d.Host,
		d.Port,
		d.User,
		d.Name,
		d.Password,
		d.SSLMode,
		d.Schema,
	)
}
