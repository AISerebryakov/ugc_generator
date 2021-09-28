package generator

import (
	"github.com/pretcat/ugc_test_task/pg"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (conf Config) GetPgConfig() pg.Config {
	return pg.Config{
		Host:     conf.Host,
		Port:     conf.Port,
		Database: conf.DbName,
		User:     conf.User,
		Password: conf.Password,
	}
}
