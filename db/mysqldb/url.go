package mysqldb

import (
	"short-url/constance"
	"short-url/models"
	"time"
)

type UrlImp interface {
	CreateUrl(shorturl string, url string, expireAt *time.Time) error
	GetUrl(shortUrl string) (*models.Url, error)
}

type Url struct {
	ShortUrl  string     `gorm:"size:20;uniqueIndex"`
	Url       string     `gorm:"size:1024" binding:"required"`
	ExpireAt  *time.Time `binding:"required"`
	CreatedAt *time.Time `binding:"required"`
}

func (mydb *mysqlDBObj) CreateUrl(shortUrl string, url string, expireAt *time.Time) error {

	createTime := time.Now()
	m := Url{
		ShortUrl:  shortUrl,
		Url:       url,
		ExpireAt:  expireAt,
		CreatedAt: &createTime,
	}
	if err := mydb.DB.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (mydb *mysqlDBObj) GetUrl(shortUrl string) (*models.Url, error) {
	u := Url{}
	result := mydb.DB.Where("short_url = ? AND expire_at> ?", shortUrl, time.Now().Format(constance.SecLocalTimeFormat)).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.Url{Url: u.Url, ShortUrl: u.ShortUrl, ExpireAt: u.ExpireAt}, nil
}
