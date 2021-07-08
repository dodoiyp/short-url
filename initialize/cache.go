package initialize

import (
	"fmt"
	"short-url/pkg/global"

	lru "github.com/hashicorp/golang-lru"
)

func Cache() {
	cache, err := lru.New(global.Conf.Cache.CacheSize)
	if err != nil {
		panic(fmt.Errorf("fatal initialize cache: %w", err))
	}
	global.Cache = cache
}
