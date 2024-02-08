package models

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Events map[string]Event `json:"events" yaml:"events"`
}

func (c *Config) ParseConfig(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	extension := getFileExtension(configFile)
	switch extension {
	case "yaml", "yml":
		err = yaml.Unmarshal(data, c)
	case "json":
		err = json.Unmarshal(data, c)
	default:
		err = errors.New("unsupported file extension")
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) ReloadConfig(configFile string) error {
	return c.ParseConfig(configFile)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
