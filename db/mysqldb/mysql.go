package mysqldb

import (
	"fmt"
	"short-url/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mysqlDbConfig *MySQLConfig
	mysqlDb       MySQLDBAccessObject
	once          sync.Once

	defaultMaxConnection  = 200
	defaultIdleConnection = 10
)

type MySQLConfig struct {
	MaxConnection  int
	IdleConnection int
	LogMode        int
	Dsn            string
	ShowDsn        string
}

func LoadConfig(cfg *config.MysqlConfiguration) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		cfg.Collation,
		cfg.Query,
	)
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		cfg.Username,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		cfg.Collation,
		cfg.Query,
	)

	mysqlDbConfig = &MySQLConfig{
		MaxConnection:  cfg.MaxConnection,
		IdleConnection: cfg.IdleConnection,
		LogMode:        cfg.LogMode,
		Dsn:            dsn,
		ShowDsn:        showDsn,
	}
	if mysqlDbConfig.MaxConnection <= 0 {
		mysqlDbConfig.MaxConnection = defaultMaxConnection
	}

	if mysqlDbConfig.IdleConnection <= 0 {
		mysqlDbConfig.IdleConnection = defaultIdleConnection
	}

}

func RetriveMySQLDBAccessObj() MySQLDBAccessObject {
	once.Do(func() {
		mysqlDb = &mysqlDBObj{}
	})
	return mysqlDb
}

func Start() error {
	var err error
	mysqlDb = RetriveMySQLDBAccessObj()
	mysqlDb, err = initMySqlDB(mysqlDbConfig)
	if err != nil {
		return err
	}

	return err
}

type mysqlDBObj struct {
	DB *gorm.DB
}

func (db *mysqlDBObj) Close() error {
	d, err := db.DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

type MySQLDBAccessObject interface {
	Close() error

	UrlImp
	SequenceImp
}

func initMySqlDB(cfg *MySQLConfig) (MySQLDBAccessObject, error) {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{}); err != nil {
		return nil, fmt.Errorf("Connection to MySQL DB error : %v", err)
	}

	// config db open & idle connection nums
	d, err := db.DB()
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(cfg.MaxConnection)
	d.SetMaxIdleConns(cfg.IdleConnection)

	if err = d.Ping(); err != nil {
		return nil, fmt.Errorf("Ping MySQL db error : %v", err)
	}

	// 在 gorm.io/gorm 後 db.LogMode(c.DBLogEnable) 改為按照 log 等級
	db.Logger.LogMode(logger.LogLevel(cfg.LogMode)) // Silent=1, Error=2, Warn=3, Info=4

	if err = db.AutoMigrate(
		&Url{},
		&Sequence{},
	); err != nil {
		return nil, err
	}

	return &mysqlDBObj{DB: db}, err
}
