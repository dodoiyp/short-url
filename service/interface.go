package service

import (
	"short-url/models"
	"time"
)

type ShortUrlSevice interface {
	CreateShortUrl(url string, expireAt *time.Time) (shorturl string, err error)
	GetUrl(shortUrl string) (*models.Url, error)
}
