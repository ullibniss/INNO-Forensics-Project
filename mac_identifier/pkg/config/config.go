package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"sync"
)

type Config struct {
	DB_NAME string `yaml:"db_name"`
	DB_USER string `yaml:"db_user"`
	DB_PASS string `yaml:"db_pass"`
	DB_PORT string `yaml:"db_port"`
	DB_HOST string `yaml:"db_host"`
}

var c *Config
var once sync.Once

func LoadConfig() *Config {

	filename := os.Getenv("CONFIG_PATH")
	if filename == "" {
		filename = "config.yml"
	}

	once.Do(func() {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var buffer []byte
		buffer, err = io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		c = &Config{}
		err = yaml.Unmarshal(buffer, c)
		if err != nil {
			log.Fatal(err)
		}
	})
	return c
}
