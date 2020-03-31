package config

import (
	"go-restful-api/app/compoments/mysql"
	"go-restful-api/app/compoments/redis"
	"gopkg.in/yaml.v3"
	"os"
)

type Options struct {
	RedisConfig *redis.RedisConfig `yaml:"redis"`
	DBConfig    *mysql.DbConfig    `yaml:"mysql"`
}

func ReadYamlConfig(path string) (*Options, error) {
	config := &Options{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(config)
	}
	return config, nil
}

func LoadConfig() (*Options, error) {
	conf, err := ReadYamlConfig("/Users/galen/helinhan/go-restful-api/config/config.yaml")
	if err != nil {
		return nil, err
	}
	return conf, nil
}
