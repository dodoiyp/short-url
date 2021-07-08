package service

type ShortUrlSevice interface {
	CreateShortUrl(u *Url) error
	GetUrl(shortUrl string) (*Url, error)
	NewKey() (string, error)
}
