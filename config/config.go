package config

import "github.com/spf13/viper"

type Config struct {
	DB   DBConfig   `mapstructure:",squash"`
	Host HostConfig `mapstructure:",squash"`
	Jwt  JwtConfig  `mapstructure:",squash"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Name     string `mapstructure:"DB_NAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
}

type HostConfig struct {
	Port         string   `mapstructure:"HOST_PORT"`
	AllowOrigins []string `mapstructure:"HOST_ALLOW_ORIGINS"`
}

type JwtConfig struct {
	Secret     string `mapstructure:"JWT_SECRET"`
	Expiration int    `mapstructure:"JWT_EXPIRATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.SetDefault("HOST_PORT", "8000")
	viper.SetDefault("HOST_ALLOW_ORIGINS", []string{"*"})
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("JWT_SECRET", "mydevsecret")
	viper.SetDefault("JWT_EXPIRATION", 604800)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return

}
