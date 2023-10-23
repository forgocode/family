package conf

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"

	"github.com/forgocode/family/internal/pkg/newlog"
)

type Config struct {
	Mysql MysqlConfig `json:"mysql" yaml:"mysql"`
	Redis RedisConfig `json:"redis" yaml:"redis"`
	Mongo MongoConfig `json:"mongo" yaml:"mongo"`
}

type MysqlConfig struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     uint16 `json:"port" yaml:"port"`
	DB       string `json:"db" yaml:"db"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

type RedisConfig struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     uint16 `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
}

type KafkaConfig struct {
}

type MongoConfig struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     uint16 `json:"port" yaml:"port"`
	DB       string `json:"db" yaml:"db"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

func GetConfig() Config {
	loadConfig.Do(load)
	return *c
}

var c *Config
var configPath = "/root/goWorkspace/src/family/internal/conf/config.yaml"

var loadConfig sync.Once

func load() {
	dataBytes, err := os.ReadFile(configPath)
	if err != nil {
		newlog.Logger.Errorf("failed to load config from:%s, err: %+v\n", configPath, err)
		return
	}
	//c := &Config{}
	err = yaml.Unmarshal(dataBytes, &c)
	if err != nil {
		newlog.Logger.Errorf("failed to parse config err: %+v\n", err)
		return
	}

}
