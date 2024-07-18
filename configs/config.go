package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	WebServerPort   string `mapstructure:"WEB_SERVER_PORT"`
	AppHost         string `mapstructure:"APP_HOST"`
	MaxRequestToken int    `mapstructure:"MAX_REQUESTS_TOKEN"`
	MaxRequestIp    int    `mapstructure:"MAX_REQUESTS_IP"`
	BlockTimeIp     int    `mapstructure:"BLOCK_TIME_IP"`
	BlockTimeToken  int    `mapstructure:"BLOCK_TIME_TOKEN"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, err
}
