package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORTGRPC    int    `mapstructure:"PORT_GRPC"`
	HOST        string `mapstructure:"HOST"`
	STORAGEDIR  string `mapstructure:"STORAGEDIR"`
	LISTLIMIT   int    `mapstructure:"LISTLIMIT"`
	UPLOADLIMIT int    `mapstructure:"UPLOADLIMIT"`
	PORTMETRICS int    `mapstructure:"PORT_METRICS"`
}

func LoadConfig(path string) (cfg *Config, err error) {

	cfg = new(Config)

	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {

		return nil, err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {

		return nil, err
	}

	return cfg, nil
}
