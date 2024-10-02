package config

import (
	"flag"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
)

var cfg *Config = &Config{}

type Config struct {
	Server      Server
	Postgres    Postgres
	Env         string
	Redis       RedisConfigs
	Hostname    string
}

func GetEnv() string {
	return cfg.Env
}

func isTestEnv() bool {
	return flag.Lookup("test.v") != nil
}

func setTestEnvVariable() {
	cfg.Env = EnvTest
}

func Load() error {
	config, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "config.New")
	}
	if err := config.Load(); err != nil {
		return errors.Wrap(err, "config.Load")
	}
	if err := config.Scan(cfg); err != nil {
		return errors.Wrap(err, "config.Scan")
	}
	if isTestEnv() {
		setTestEnvVariable()
	}
	return nil
}
