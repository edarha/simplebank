package util

import "github.com/spf13/viper"

var Conf *Config

type (
	Config struct {
		Server   Server   `mapstructure:"server"`
		Postgres Postgres `mapstructure:"POSTGRES"`
	}

	Server struct {
		Address string `mapstructure:"address"`
	}

	Postgres struct {
		Driver string `mapstructure:"driver"`
		Source string `mapstructure:"source"`
	}
)

func LoadConfig(path string) (*Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
