package service

import (
	"short-url/pkg/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const timeFormat = "2006-01-02 15:04:05"

type Url struct {
	ShortUrl  string    `gorm:"size:20;uniqueIndex"`
	Url       string    `gorm:"size:1024" binding:"required"`
	ExpireAt  time.Time `binding:"required"`
	CreatedAt time.Time `binding:"required"`
}

type Sequence struct {
	ID int `gorm:"size:20;uniqueIndex"`
}

func NewShortService(c *gin.Context) ShortUrlSevice {
	return &Service{
		db: global.Mysql,
	}
}

type Service struct {
	db *gorm.DB
}

func (s *Service) CreateShortUrl(u *Url) error {
	if err := s.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

//TODO : add where expiretime filter
func (s *Service) GetUrl(shortUrl string) (*Url, error) {
	u := &Url{}
	result := s.db.Where("short_url = ? AND expire_at> ?", shortUrl, time.Now().Format(timeFormat)).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
func (s *Service) NewKey() (string, error) {
	seq := Sequence{}
	err := s.db.Create(&seq).Error
	if err != nil {
		return "", err
	}
	id := strconv.FormatInt(int64(seq.ID), 10)
	return id, nil
}
