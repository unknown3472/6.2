package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadJSONConfig(t *testing.T) {
    type Config struct {
        Port int `json:"port"`
    }

    var config Config
    err := LoadJSONConfig("testdata/config.json", &config)
    if err != nil {
        t.Errorf("Failed to load JSON config: %v", err)
    }

    if config.Port != 8080 {
        t.Errorf("Expected Port to be 8080, got %d", config.Port)
    }
}
func TestLoadYAMLConfig(t *testing.T) {
    type Config struct {
        Port int `yaml:"port"`
    }

    var config Config
    err := LoadYAMLConfig("testdata/config.yaml", &config)
    if err != nil {
        t.Errorf("Failed to load YAML config: %v", err)
    }

    if config.Port != 8080 {
        t.Errorf("Expected Port to be 8080, got %d", config.Port)
    }
}

func TestLoadEnvConfig(t *testing.T) {
	godotenv.Load()
    err := LoadEnvConfig(".env")
    if err != nil {
        t.Errorf("Failed to load ENV config: %v", err)
    }

    if os.Getenv("PORT") != "8080" {
        t.Errorf("Expected PORT to be 8080, got %s", os.Getenv("PORT"))
    }
}
