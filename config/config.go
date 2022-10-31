package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Postgres struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		DBName  string `yaml:"db-name"`
		SSLmode string `yaml:"sslmode"`
	} `yaml:"postgres"`
}

func GetConfig(path string) *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		help, _ := cleanenv.GetDescription(cfg, nil)
		log.Fatalln(err, help)
	}
	return cfg
}
