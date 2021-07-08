package initialize

import (
	"fmt"
	"short-url/pkg/global"

	"github.com/spf13/viper"
)

func Config() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&global.Conf)
	if err != nil {
		panic(fmt.Errorf("fatal error config unmarshal: %w", err))
	}
}
