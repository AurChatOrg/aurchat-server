package config

import (
	"bufio"
	"os"
	"strings"
)

// Read policy
const (
	StrategyEnv  = "ENV_VAR" // Read from environment variables
	StrategyFile = "FILE"    // Read from .env file
)

// readEnv
func readEnv(key, defaultValue string) string {
	if strategy == StrategyEnv { // If the read strategy is StrategyEnv, then retrieve the value from the environment variable
		return getEnv(key, defaultValue)
	}

	return dotEnvCache[key] // If the read strategy is StrategyFile, then directly return the content in the cache
}

// getEnv Read environment variables
func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	} else {
		return defaultValue
	}
}

// .env file cache
var dotEnvCache = map[string]string{}
var strategy string

func init() {
	// Try reading the ENV field in .env
	parseDotEnv()
	strategy = dotEnvCache["ENV"]
	if strategy == "" {
		strategy = StrategyEnv // Default reading of environment variables
	}

	// If the strategy is FILE, parseDotEnv() will automatically load the. env file into the cache
}

// parseDotEnv Simple KV analysis
func parseDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		dotEnvCache[key] = val
	}
}
