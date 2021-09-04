package config

import "github.com/spf13/viper"

type Config struct {
	DB   DBConfig   `mapstructure:",squash"`
	Host HostConfig `mapstructure:",squash"`
}

type DBConfig struct {
	Url string `mapstructure:"DB_URL"`
}

type HostConfig struct {
	Port         string   `mapstructure:"HOST_PORT"`
	AllowOrigins []string `mapstructure:"HOST_ALLOW_ORIGINS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.SetDefault("HOST_PORT", "8000")
	viper.SetDefault("HOST_ALLOW_ORIGINS", []string{"*"})

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return

}
