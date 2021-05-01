package db

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	SSLMode      string
	Timezone     string
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.Username,
		c.Password,
		c.DatabaseName,
		c.Port,
		c.SSLMode,
		c.Timezone,
	)
}

func NewConfigFromEnv() (*Config, error) {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	var iPort int
	if port == "" {
		iPort = 5432
	} else {
		var err error
		iPort, err = strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
	}
	sslmode := os.Getenv("DB_SSL_MODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	timezone := os.Getenv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "Europe/Sofia"
	}

	return &Config{
		Host:         host,
		Username:     username,
		Password:     password,
		DatabaseName: name,
		Port:         iPort,
		SSLMode:      sslmode,
		Timezone:     timezone,
	}, nil
}
