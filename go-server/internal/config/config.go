package config

import "time"

// Config is the main app configuration
type Config struct {
	// Port is the http server port
	Port string
	// Host is the http server host
	Host string
	// ShutdownTimeout is the amount of seconds to wait before forcing http server shutdown
	ShutdownTimeout time.Duration

	// JWTSecret is the secret used to encrypt jwt tokens
	JWTSecret string
}

func NewConfig(getevn func(string) string) *Config {
	// Default values
	cfg := &Config{
		Port:            "8888",
		Host:            "0.0.0.0",
		ShutdownTimeout: time.Duration(5),
		JWTSecret:       "secret",
	}

	if len(getevn("jwtsecret")) > 0 {
		cfg.JWTSecret = getevn("jwtsecret")
	}

	if len(getevn("port")) > 0 {
		cfg.Port = getevn("port")
	}

	if len(getevn("host")) > 0 {
		cfg.Port = getevn("host")
	}

	if len(getevn("shutdowntimeout")) > 0 {
		cfg.Port = getevn("shutdowntimeout")
	}

	return cfg
}
