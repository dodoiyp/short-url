package local_cache

import (
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

var (
	localCacheConfig *LocalCacheConfig
	loaclCache       *LocalCacheObject
	once             sync.Once
)

type LocalCacheObject struct {
	cache *lru.Cache
}
type LocalCacheConfig struct {
	CacheSize int
}

func LoadCacheConfig(cacheSize int) {
	if cacheSize <= 0 {
		cacheSize = 0
	}
	localCacheConfig = &LocalCacheConfig{CacheSize: cacheSize}
}

func RetriveLocalCacheAccessObj() *LocalCacheObject {
	once.Do(func() {
		loaclCache = &LocalCacheObject{}
	})
	return loaclCache
}

func Start() error {
	var err error
	loaclCache = RetriveLocalCacheAccessObj()
	loaclCache, err = initLocalCache(localCacheConfig)
	if err != nil {
		return err
	}
	return err
}

func initLocalCache(cfg *LocalCacheConfig) (*LocalCacheObject, error) {
	cache, err := lru.New(cfg.CacheSize)
	if err != nil {
		return nil, err
	}
	return &LocalCacheObject{cache: cache}, nil
}
