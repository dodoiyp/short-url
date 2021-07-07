package global

type Configuration struct {
	System    SystemConfiguration    `mapstructure:"system" json:"system"`
	Mysql     MysqlConfiguration     `mapstructure:"mysql" json:"mysql"`
	RateLimit RateLimitConfiguration `mapstructure:"rate-limit" json:"rateLimit"`
}

type SystemConfiguration struct {
	ApiVersion     string `mapstructure:"api-version" json:"apiVersion"`
	Port           int    `mapstructure:"port" json:"port"`
	PprofPort      int    `mapstructure:"pprof-port" json:"pprofPort"`
	ConnectTimeout int    `mapstructure:"connect-timeout" json:"connectTimeout"`
}

type MysqlConfiguration struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type RateLimitConfiguration struct {
	Max int64 `mapstructure:"max" json:"max"`
}
