package cache

import (
	"short-url/cache/local_cache"
	"short-url/config"
	"time"
)

type CacheImpement interface {
	Find(shortUrl string) (string, bool)
	Insert(shortUrl string, url string, expireAt *time.Time) error
}

func RetriveCacheAccessModel() CacheImpement {
	return local_cache.RetriveLocalCacheAccessObj()
}

//TODO add external cache and nil check
func InitCache(cfg *config.CacheConfiguration) error {
	return initLocalCache(&cfg.LocalCacheConfiguration)
}

func initLocalCache(cfg *config.LocalCacheConfiguration) error {
	local_cache.LoadCacheConfig(cfg.CacheSize)
	return local_cache.Start()
}
