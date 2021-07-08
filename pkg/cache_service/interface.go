package cache_service

type CacheService interface {
	Find(shortUrl string) (string, bool)
	Insert(shortUrl string, url *Url) error
}
