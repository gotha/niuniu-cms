package main

import (
	"os"
	"strconv"
)

const defaultPort = 8080

type Config struct {
	Port int
}

func NewConfigFromEnv() (*Config, error) {
	var iPort int
	port := os.Getenv("PORT")
	if port == "" {
		iPort = defaultPort
	} else {
		var err error
		iPort, err = strconv.Atoi(port)
		if err != nil {
			return nil, err
		}
	}

	return &Config{
		Port: iPort,
	}, nil
}
