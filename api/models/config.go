package models

import (
	"os"

	"github.com/spf13/viper"
)

// Config ...
var (
	Config *ConfigSetup
)

func init() {
	InitConfig()
}

// LoadConfig ...
func LoadConfig(file string) {
	if Config != nil {
		return
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	Config = new(ConfigSetup)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&Config)
}

// InitConfig ...
func InitConfig() {
	LoadConfig("api/app.yaml")
}

// ConfigSetup ...
type ConfigSetup struct {
	BlockLevelConfig BlockLevelConfig `yaml:"BlockLevelConfig"`
}

type BlockLevelConfig struct {
	Level string `yaml:"Level"`
}
