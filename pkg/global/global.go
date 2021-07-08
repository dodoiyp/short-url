package global

import (
	lru "github.com/hashicorp/golang-lru"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Conf  Configuration
	Mysql *gorm.DB
	Cache *lru.Cache
	Log   *zap.SugaredLogger
)
