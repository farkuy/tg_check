package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ConnectTime  string `yaml:"connect_timeout"`
	ReadTime     string `yaml:"read_timeout"`
	SenTime      string `yaml:"send_timeout"`
	DatabasePath string `yaml:"database_path"`
	Server
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

const env = "../../.env"

func Init() (*Config, error) {
	err := godotenv.Load(env)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения .env файла: %v", err)
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return nil, fmt.Errorf("Путь конфига не найден: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения config.yaml: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создание конфига: %v", err)
	}

	return &config, nil
}
