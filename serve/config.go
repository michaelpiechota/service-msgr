package serve

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

const configPath = "."
const configType = "env"
const defaultConfigName = ".env.template"

func NewConfig(configFileName string) *Config {
	if configFileName == "" {
		configFileName = defaultConfigName
	}

	v := viper.New()
	v.AutomaticEnv()
	v.AddConfigPath(configPath)
	v.SetConfigType(configType)
	v.SetConfigName(configFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{Viper: v}
}

// GetStringOrPanic returns a config value as a string.
// It panics if the key does not exist in the config layer.
func (c *Config) GetStringOrPanic(key string) string {
	if v := c.Viper.GetString(key); v != "" {
		return v
	}

	panic(fmt.Sprintf("environment variable %s not found", key))
}

// GetStringOrDefault returns a config value as a string.
// It returns a default if the key does not exist in the config layer.
func (c *Config) GetStringOrDefault(key string, def string) string {
	if v := c.Viper.GetString(key); v != "" {
		return v
	}

	return def
}
