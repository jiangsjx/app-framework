package kit

import "github.com/spf13/viper"

const DefaultConfig = "config.yml"

type Config struct {
	App struct {
		Host string
		Port string
	}

	Log struct {
		Debug bool
		File  string
	}

	MySQL struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Database int
	}
}

func NewConfig(file string) (*Config, error) {
	var config Config

	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
