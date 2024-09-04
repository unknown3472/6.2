package config

import (
	"encoding/json"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// JSON konfiguratsiyani yuklash funksiyasi
func LoadJSONConfig(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		return err
	}

	return nil
}

// YAML konfiguratsiyani yuklash funksiyasi
func LoadYAMLConfig(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return err
	}

	return nil
}

// ENV konfiguratsiyani yuklash funksiyasi
func LoadEnvConfig(envFilePath string) error {
	file, err := os.Open(envFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
