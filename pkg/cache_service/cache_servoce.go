package cache_service

import (
	"fmt"
	"short-url/pkg/global"
	"time"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru"
)

type LocalCacheService struct {
	cache *lru.Cache
	//redis *redis.QueryRedis
}
type Url struct {
	Url      string    `gorm:"size:1024" binding:"required"`
	ExpireAt time.Time `binding:"required"`
}

func New(c *gin.Context) CacheService {
	return &LocalCacheService{
		cache: global.Cache,
	}
}

func (lcs *LocalCacheService) Find(shortUrl string) (string, bool) {
	val, found := lcs.cache.Get(shortUrl)
	if !found {
		return "", !found
	}

	url, ok := val.(Url)
	if !ok {
		fmt.Printf("cache data unsupport key:%v , data %v ", shortUrl, val)
	}

	if url.ExpireAt.Before(time.Now()) {
		return "", !found
	}

	return url.Url, found
}

func (lcs *LocalCacheService) Insert(shortUrl string, url *Url) error {
	lcs.cache.Add(shortUrl, url)
	return nil
}
