package config

import (
	"sync"
)

var (
	once sync.Once
	cfg  *Config
	Cfg  *Config
)

// Load Load env
func Load() *Config {
	once.Do(func() {
		cfg = LoadYAMLConfig()
	})

	Cfg = cfg
	return cfg
}
