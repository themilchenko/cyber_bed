package config

import (
	"flag"
	"fmt"
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
	Database struct {
		User    string `yaml:"postgres"`
		DbName  string `yaml:"cyber_garden"`
		Host    string `yaml:"localhost"`
		Port    uint64 `yaml:"port"`
		SslMode string `yaml:"sslmode"`
	} `yaml:"database"`
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
		Database: struct {
			User    string `yaml:"postgres"`
			DbName  string `yaml:"cyber_garden"`
			Host    string `yaml:"localhost"`
			Port    uint64 `yaml:"port"`
			SslMode string `yaml:"sslmode"`
		}(struct {
			User    string
			DbName  string
			Host    string
			Port    uint64
			SslMode string
		}{
			User:    "postgres",
			DbName:  "cyber_garden",
			Host:    "localhost",
			Port:    5432,
			SslMode: "disable",
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

func (c *Config) FormatDbAddr() string {
	return fmt.Sprintf(
		"user=%s dbname=%s host=%s port=%d sslmode=%s",
		c.Database.User,
		c.Database.DbName,
		c.Database.Host,
		c.Database.Port,
		c.Database.SslMode,
	)
}

func ParseFlag(path *string) {
	flag.StringVar(path, "ConfigPath", "configs/setup.yaml", "Path to Config")
}
