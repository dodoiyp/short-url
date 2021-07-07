package global

import "database/sql"

var (
	Conf  Configuration
	Mysql *sql.DB
)
