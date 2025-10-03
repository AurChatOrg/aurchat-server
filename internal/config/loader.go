package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// LoadYAMLConfig Load configuration from YAML file
func LoadYAMLConfig() *Config {
	configPath := getEnv("CONFIG_PATH", "./config.yaml")

	// If the configuration file does not exist, use the built-in default configuration
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return getDefaultConfig()
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		// The file exists but failed to read, reverting to default configuration
		return getDefaultConfig()
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		// YAML parsing failed, reverting to default configuration
		return getDefaultConfig()
	}

	return &config
}

// getDefaultConfig Provide built-in default configuration
func getDefaultConfig() *Config {
	return &Config{
		App: App{
			Name: "aurchat-gateway",
			Env:  "dev",
		},
		HTTP: HTTP{
			Listen: ":8080",
		},
		Redis: Redis{
			Addr: "127.0.0.1:6379",
			DB:   0,
		},
		NATS: NATS{
			URL: "nats://127.0.0.1:4222",
		},
		Database: Database{
			DSN: "host=localhost user=user password=pass dbname=aurchat port=5432 sslmode=disable",
		},
		Auth: Auth{
			Keys: "abcd1234abcd1234abcd1234abcd1234",
			TTL:  259200,
		},
		Snowflake: Snowflake{
			WorkerID:          1,
			WorkerIDBitLength: 8,
			SeqBitLength:      12,
			BaseTime:          "2020-01-01",
		},
	}
}

// LoadEnvConfigs Load configuration from environment variables
func LoadEnvConfigs(baseConfig *Config) *Config {
	config := *baseConfig

	// App Configuration
	if v := os.Getenv("APP_NAME"); v != "" {
		config.App.Name = v
	}
	if v := os.Getenv("APP_ENV"); v != "" {
		config.App.Env = v
	}

	// HTTP Configuration
	if v := os.Getenv("HTTP_LISTEN"); v != "" {
		config.HTTP.Listen = v
	}

	// Redis Configuration
	if v := os.Getenv("REDIS_ADDR"); v != "" {
		config.Redis.Addr = v
	}
	if v := os.Getenv("REDIS_DB"); v != "" {
		if db, err := strconv.Atoi(v); err == nil {
			config.Redis.DB = db
		}
	}

	// NATS Configuration
	if v := os.Getenv("NATS_URL"); v != "" {
		config.NATS.URL = v
	}

	// Database Configuration
	if v := os.Getenv("POSTGRES_DSN"); v != "" {
		config.Database.DSN = v
	}

	// Authentication Configuration
	if v := os.Getenv("AUTH_KEY"); v != "" {
		config.Auth.Keys = v
	}
	if v := os.Getenv("AUTH_TTL"); v != "" {
		if ttl, err := strconv.Atoi(v); err == nil {
			config.Auth.TTL = uint32(ttl)
		}
	}

	// Snowflake Configuration
	if v := os.Getenv("SNOWFLAKE_WORKER_ID"); v != "" {
		if workerID, err := strconv.Atoi(v); err == nil {
			config.Snowflake.WorkerID = int64(workerID)
		}
	}

	return &config
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
