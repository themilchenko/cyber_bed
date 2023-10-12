package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	address     = "localhost"
	port        = 8080
	loggerLevel = "debug"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Port    uint64 `yaml:"port"`
	} `yaml:"server"`
	LoggerLvl string `yaml:"logger_level"`
}

func New() *Config {
	return &Config{
		Server: struct {
			Address string `yaml:"address"`
			Port    uint64 `yaml:"port"`
		}(struct {
			Address string
			Port    uint64
		}{
			Address: address,
			Port:    port,
		}),
		LoggerLvl: loggerLevel,
	}
}

func (c *Config) Open(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = yaml.NewDecoder(file).Decode(c); err != nil {
		return err
	}

	return nil
}

func ParseFlag(path *string) {
	flag.StringVar(path, "ConfigPath", "configs/setup.yaml", "Path to Config")
}
