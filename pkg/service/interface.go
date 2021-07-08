package service

import "short-url/models"

type ShortUrlSevice interface {
	CreateShortUrl(u *models.Url) error
	GetUrl(shortUrl string) (*models.Url, error)
	NewKey() (string, error)
}
