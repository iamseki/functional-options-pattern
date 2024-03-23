package main

import (
	"errors"
	"log"
)

type DBConfigOptions struct {
	Port              int
	PoolSize          int
	ConnectionTimeout int
}

type Option func(*DBConfigOptions)

func WithPort(port int) Option {
	return func(c *DBConfigOptions) {
		c.Port = port
	}
}

func WithPoolSize(size int) Option {
	return func(c *DBConfigOptions) {
		c.PoolSize = size
	}
}

func WithConnectionTimeout(timeout int) Option {
	return func(c *DBConfigOptions) {
		c.ConnectionTimeout = timeout
	}
}

// It's a handy and API-friendly way to handle options, consumers can opt for default values and doesn't pass anything for this init function.
func InitFakeDB(connectionURI string, opts ...Option) error {
	cfg := &DBConfigOptions{}
	for _, opt := range opts {
		opt(cfg)
	}

	if connectionURI == "" {
		return errors.New("mandatory config is nil, please provided one connectionURI")
	}

	log.Printf("init with cfg options: %v \n", cfg)
	// We can also set defaults for PoolSize and other optional parameters...
	return nil
}

func main() {
	// Here's how we can use the function options pattern, to call the init function WITH or WITHOUT options:
	err := InitFakeDB("")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	// With default values
	InitFakeDB("test:test@test.com")

	// Set some options
	InitFakeDB("test:test@test.com",
		WithPort(12),
		WithPoolSize(10))
}
