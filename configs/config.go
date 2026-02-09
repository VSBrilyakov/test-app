package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Config struct {
	Server ServerConfig
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}

	var loadedConfig Config
	err = yaml.Unmarshal(data, &loadedConfig)
	if err != nil {
		return nil, err
	}

	return &loadedConfig, nil
}

func (c *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
