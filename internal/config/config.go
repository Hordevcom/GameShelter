package config

import (
	"flag"
	"fmt"

	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ServerAdress string `env:"RUN_ADDRESS"`
	DatabaseDsn  string `env:"DATABASE_URI"`
	Logger       *logging.Logger
}

func NewConfig(logger *logging.Logger) Config {
	var conf Config
	err := env.Parse(&conf)

	if err != nil {
		panic(fmt.Sprintf("config parse error: %v", err))
	}

	if conf.DatabaseDsn != "" && conf.ServerAdress != "" {
		fmt.Printf("env: %v", conf.DatabaseDsn)
		return conf
	}

	if conf.DatabaseDsn == "" {
		flag.StringVar(&conf.DatabaseDsn, "d", "", "database dsn") //"postgres://postgres:1@localhost:5432/postgres"
	}

	flag.StringVar(&conf.ServerAdress, "a", "localhost:8080", "server address")

	flag.Parse()

	return conf
}
