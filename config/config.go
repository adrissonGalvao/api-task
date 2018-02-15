package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("/home/adrisson/go/src/api-task/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
