package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"prod"`
	GRPC GRPCConfig `yaml:"grpc" env-required:"true"`
	DB   Storage    `yaml:"db" env-required:"true"`
}

type Storage struct {
	DbUser  string `yaml:"db_user" env-required:"true"`
	DbPass  string `yaml:"db_pass" env-required:"true"`
	DbName  string `yaml:"db_name" env-required:"true"`
	SslMode string `yaml:"ssl_mode" env-required:"true"`
	Port    string `yaml:"port" env-default:"5432"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(path string) *Config {
	if path == "" {
		panic("config path os empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exists: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

func FetchConfigPath() string {
	var res string

	// --config="path/..."
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("config")
	}

	// CONFIG_PATH=./path/to/config/file.yaml sso
	// sso --config=./path...
	return res
}
