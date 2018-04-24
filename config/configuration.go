package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

type Configuration interface {
	GetString(key string) string
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.SetEnvPrefix(`nats-rest`)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func NewViperConfig() Configuration {
	v := &viperConfig{}
	v.Init()
	return v
}
