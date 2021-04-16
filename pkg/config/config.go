package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func Get() *Config {
	f, err := os.Open("configs/config.yml")
	if err != nil {
		os.Exit(2)
		return nil
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)

	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		os.Exit(3)
		return nil
	}

	return &config
}