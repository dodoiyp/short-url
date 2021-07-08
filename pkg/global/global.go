package global

import (
	lru "github.com/hashicorp/golang-lru"
	"gorm.io/gorm"
)

var (
	Conf  Configuration
	Mysql *gorm.DB
	Cache *lru.Cache
)
