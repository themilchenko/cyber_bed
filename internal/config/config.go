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
		User    string `yaml:"user"`
		DbName  string `yaml:"dbname"`
		Host    string `yaml:"host"`
		Port    uint64 `yaml:"port"`
		SslMode string `yaml:"sslmode"`
	} `yaml:"database"`
	LoggerLvl    string `yaml:"logger_level"`
	RecognizeAPI struct {
		MaxImages    int    `yaml:"max_images"`
		BaseURL      string `yaml:"base_url"`
		CountResults int    `yaml:"count_results"`
		ImageField   string `yaml:"image_field"`
		Token        string `yaml:"token"`
	} `yaml:"recognize_api"`
	TrefleAPI struct {
		BaseURL     string `yaml:"base_url"`
		CountPlants int    `yaml:"count_plants"`
		Token       string `yaml:"token"`
	} `yaml:"trefle_api"`
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
			User    string `yaml:"user"`
			DbName  string `yaml:"dbname"`
			Host    string `yaml:"host"`
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
		RecognizeAPI: struct {
			MaxImages    int    `yaml:"max_images"`
			BaseURL      string `yaml:"base_url"`
			CountResults int    `yaml:"count_results"`
			ImageField   string `yaml:"image_field"`
			Token        string `yaml:"token"`
		}(struct {
			MaxImages    int
			BaseURL      string
			CountResults int
			ImageField   string
			Token        string
		}{
			MaxImages:    5,
			BaseURL:      "https://my-api.plantnet.org/v2/identify/",
			CountResults: 4,
			ImageField:   "images[]",
			Token:        "token",
		}),
		TrefleAPI: struct {
			BaseURL     string `yaml:"base_url"`
			CountPlants int    `yaml:"count_plants"`
			Token       string `yaml:"token"`
		}(struct {
			BaseURL     string
			CountPlants int
			Token       string
		}{
			BaseURL:     "https://{defaultHost}/api/v1/plants/",
			CountPlants: 5,
			Token:       "token",
		}),
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
