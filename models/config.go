package models

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Events map[string]Event `yaml:"events"`
}

func (c *Config) ParseYAML(yamlFile string) {
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (c *Config) ReloadYAML(yamlFile string) error {
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}
