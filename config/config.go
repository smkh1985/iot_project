package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type SensorConfig struct {
	ID       string        `yaml:"id"`
	Type     string        `yaml:"type"`
	Interval time.Duration `yaml:"interval"`
}

type Config struct {
	Sensors []SensorConfig `yaml:"sensors"`
}

func LoadConfig(path string) Config {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return cfg
}
