package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf *Configuration

type Configuration struct {
	System    SystemConfiguration    `mapstructure:"system" json:"system"`
	Mysql     MysqlConfiguration     `mapstructure:"mysql" json:"mysql"`
	RateLimit RateLimitConfiguration `mapstructure:"rate-limit" json:"rateLimit"`
	Logs      LogsConfiguration      `mapstructure:"logs" json:"logs"`
	Cache     CacheConfiguration     `mapstructure:"cache" json:"cache"`
}

type SystemConfiguration struct {
	Port           int `mapstructure:"port" json:"port"`
	ConnectTimeout int `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type LogsConfiguration struct {
	Level      string `mapstructure:"level" json:"level"`
	Path       string `mapstructure:"path" json:"path"`
	MaxSize    int    `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int    `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int    `mapstructure:"max-age" json:"maxAge"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
}

type MysqlConfiguration struct {
	Username       string `mapstructure:"username" json:"username"`
	Password       string `mapstructure:"password" json:"password"`
	Database       string `mapstructure:"database" json:"database"`
	Host           string `mapstructure:"host" json:"host"`
	Port           int    `mapstructure:"port" json:"port"`
	LogMode        int    `mapstructure:"log-mode" json:"logMode"`
	Charset        string `mapstructure:"charset" json:"charset"`
	Query          string `mapstructure:"query" json:"query"`
	MaxConnection  int    `mapstructure:"max-connection" json:"max-connection"`
	IdleConnection int    `mapstructure:"idle-connection" json:"idle-connection"`
	Collation      string `mapstructure:"collation" json:"collation"`
}

type RateLimitConfiguration struct {
	Max int64 `mapstructure:"max" json:"max"`
}

type CacheConfiguration struct {
	LocalCacheConfiguration `mapstructure:"local-cache" json:"local-cache"`
}
type LocalCacheConfiguration struct {
	CacheSize int `mapstructure:"cache-size" json:"cache-size"`
}

//TODO Add paremeter to load config
func InitConfig() (*Configuration, error) {
	Conf = new(Configuration)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(fmt.Errorf("fatal error config unmarshal: %w", err))
	}
	return Conf, nil
}
