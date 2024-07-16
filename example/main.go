package main

import (
	"fmt"
	"log"

	"github.com/syntaqx/env"
)

type RedisMode string

const (
	RedisModeStandalone RedisMode = "standalone"
	RedisModeCluster    RedisMode = "cluster"
)

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     int    `env:"DATABASE_PORT|DB_PORT,fallback=3306"`
	Username string `env:"DATABASE_USERNAME,default=root"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Database string `env:"DATABASE_NAME"`
}

type Config struct {
	Debug     bool           `env:"DEBUG"`
	Port      string         `env:"PORT,default=8080"`
	RedisHost []string       `env:"REDIS_HOST|REDIS_HOSTS,default=localhost:6379"`
	RedisMode RedisMode      `env:"REDIS_MODE,default=standalone"`
	Database  DatabaseConfig `env:""`
}

func main() {
	var cfg Config

	// Set example environment variables
	_ = env.Set("DEBUG", "true")
	_ = env.Set("PORT", "9090")
	_ = env.Set("REDIS_HOST", "host1,host2")
	_ = env.Set("REDIS_MODE", "cluster")
	_ = env.Set("DATABASE_HOST", "dbhost")
	_ = env.Set("DATABASE_PORT", "5432")
	_ = env.Set("DATABASE_USERNAME", "admin")
	_ = env.Set("DATABASE_PASSWORD", "secret")
	_ = env.Set("DATABASE_NAME", "mydb")

	if err := env.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
