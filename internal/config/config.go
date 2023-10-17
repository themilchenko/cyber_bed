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
	LoggerLvl      string `yaml:"logger_level"`
	CookieSettings CookieSettings
}

type CookieSettings struct {
	Secure     bool `yaml:"secure"`
	HttpOnly   bool `yaml:"http_only"`
	ExpireDate struct {
		Years  uint64 `yaml:"years"`
		Months uint64 `yaml:"months"`
		Days   uint64 `yaml:"days"`
	} `yaml:"expire_date"`
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
		CookieSettings: struct {
			Secure     bool `yaml:"secure"`
			HttpOnly   bool `yaml:"http_only"`
			ExpireDate struct {
				Years  uint64 `yaml:"years"`
				Months uint64 `yaml:"months"`
				Days   uint64 `yaml:"days"`
			} `yaml:"expire_date"`
		}(struct {
			Secure     bool
			HttpOnly   bool
			ExpireDate struct {
				Years  uint64
				Months uint64
				Days   uint64
			}
		}{
			Secure:   true,
			HttpOnly: true,
			ExpireDate: struct {
				Years  uint64
				Months uint64
				Days   uint64
			}{
				Years:  0,
				Months: 0,
				Days:   7,
			},
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
