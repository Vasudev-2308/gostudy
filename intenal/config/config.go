package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string
}
type Config struct {
	//go get -u github.com/ilyakaznacheev/cleanenv
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env_required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path to COnfig File")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config Path is not set")
		}
	}

	if _, error := os.Stat(configPath); os.IsNotExist(error) {
		log.Fatalf("Config File Does Not Exist %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Cannot Read Config File %s", err.Error())
	}

	return &cfg
}
