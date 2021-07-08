package global

import "go.uber.org/zap/zapcore"

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
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type MysqlConfiguration struct {
	Username  string `mapstructure:"username" json:"username"`
	Password  string `mapstructure:"password" json:"password"`
	Database  string `mapstructure:"database" json:"database"`
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Query     string `mapstructure:"query" json:"query"`
	LogMode   bool   `mapstructure:"log-mode" json:"logMode"`
	Charset   string `mapstructure:"charset" json:"charset"`
	Collation string `mapstructure:"collation" json:"collation"`
}

type RateLimitConfiguration struct {
	Max int64 `mapstructure:"max" json:"max"`
}

type CacheConfiguration struct {
	CacheSize int `mapstructure:"cachesize" json:"cachesize"`
}
