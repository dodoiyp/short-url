package db

import (
	"short-url/config"
	"short-url/db/mysqldb"
	"short-url/models"
	"time"
)

type DBImpement interface {
	Close() error
	CreateSequenceAndGetID() (sequenceID int, err error)
	CreateUrl(shorturl string, url string, expireAt *time.Time) error
	GetUrl(shorturl string) (url *models.Url, err error)
}

func RetriveDBAccessModel() DBImpement {
	return mysqldb.RetriveMySQLDBAccessObj()
}

func InitDataBase(cfg *config.MysqlConfiguration) error {
	return initMysql(cfg)
}

func initMysql(cfg *config.MysqlConfiguration) error {
	mysqldb.LoadConfig(cfg)
	return mysqldb.Start()
}
