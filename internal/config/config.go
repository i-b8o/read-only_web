package config

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Hook struct {
		Username string `yaml:"name"`
		Token    string `yaml:"token"`
		ChatID   string `yaml:"id"`
	} `yaml:"hook"`
	HTTP struct {
		IP   string `yaml:"ip" env:"GRPC-IP"`
		Port int    `yaml:"port" env:"GRPC-PORT"`
		CORS struct {
			AllowedMethods []string `yaml:"allowed_methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins []string `yaml:"allowed_origins"`
			AllowedHeaders []string `yaml:"allowed_headers"`
		} `yaml:"cors"`
	} `yaml:"http"`
	AppConfig struct {
		LogLevel string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
	} `yaml:"app"`
	Template struct {
		Path string `yaml:"path" env:"TEMPLATEPATH"`
	} `yaml:"template"`
	Certs struct {
		Path string `yaml:"path"`
	} `yaml:"certs"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"POSTGRES_USER" env-required:"true"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"POSTGRES_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"POSTGRES_PORT" env-required:"true"`
		Database string `yaml:"database" env:"POSTGRES_DB" env-required:"true"`
	} `yaml:"postgresql"`
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		curdir, _ := os.Getwd()
		flag.StringVar(&configPath, FlagConfigPathName, curdir+"/.configs/config.local.yaml", "this is app config file")
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "Read Only"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
