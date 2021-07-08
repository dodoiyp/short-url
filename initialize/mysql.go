package initialize

import (
	"context"
	"fmt"
	"log"
	"short-url/pkg/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// init myql
func Mysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database,
		global.Conf.Mysql.Charset,
		global.Conf.Mysql.Collation,
		global.Conf.Mysql.Query,
	)
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database,
		global.Conf.Mysql.Charset,
		global.Conf.Mysql.Collation,
		global.Conf.Mysql.Query,
	)
	log.Printf("db SDN: %s", showDsn)
	init := false
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(global.Conf.System.ConnectTimeout)*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("mysql connection timeout %v", global.Conf.System.ConnectTimeout))
				}
				return
			}
		}
	}()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("mysql init error : %v", err))
	}
	init = true

	global.Mysql = db
	log.Printf("mysql init done")
}
