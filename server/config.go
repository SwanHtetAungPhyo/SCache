package server

import (
	"net"
	"time"
)
type Option func(*Config)

type Config struct {
	Port string
	Capcity int
	DefaultExpire time.Duration
	Listener net.Listener
}

func NewCofig(opts ...Option)  (*Config, error){
	config := &Config{
		Port: ":5000",
		Capcity: 100,
		DefaultExpire: 5 * time.Minute,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config, nil
}

func WithPort(port string) Option{
	return func(c *Config) {
		c.Port = port
	}
}

func WithCapacity(capacity int) Option{
	return func(c *Config) {
		c.Capcity = capacity
	}
}

func WithExpiration(expiration time.Duration) Option {
	return func(c *Config) {
		c.DefaultExpire = expiration
	}
}