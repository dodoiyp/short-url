package service

import (
	"short-url/models"
	"short-url/pkg/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewShortService(c *gin.Context) ShortUrlSevice {
	return &Service{
		db: global.Mysql,
	}
}

type Service struct {
	db *gorm.DB
}

func (s *Service) CreateShortUrl(u *models.Url) error {
	if err := s.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUrl(shortUrl string) (*models.Url, error) {
	u := &models.Url{}
	result := s.db.Where("short_url = ? AND expire_at> ?", shortUrl, time.Now().Format(global.SecLocalTimeFormat)).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
func (s *Service) NewKey() (string, error) {
	seq := models.Sequence{}
	err := s.db.Create(&seq).Error
	if err != nil {
		return "", err
	}
	id := strconv.FormatInt(int64(seq.ID), 10)
	return id, nil
}
