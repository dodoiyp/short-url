package local_cache

import (
	"fmt"
	"time"
)

type Url struct {
	Url      string
	ExpireAt *time.Time
}

func (lco *LocalCacheObject) Find(shortUrl string) (string, bool) {
	val, found := lco.cache.Get(shortUrl)
	if !found {
		return "", !found
	}

	url, ok := val.(*Url)
	if !ok {
		fmt.Printf("cache data unsupport key:%v , data %v ", shortUrl, val)
		return "", !found
	}

	if url.ExpireAt.Before(time.Now()) {
		return "", !found
	}

	return url.Url, found
}

func (lco *LocalCacheObject) Insert(shortUrl string, url string, expireAt *time.Time) error {
	u := &Url{
		Url:      url,
		ExpireAt: expireAt,
	}
	lco.cache.Add(shortUrl, u)
	return nil
}
