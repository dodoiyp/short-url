package service

import (
	"short-url/models"
	"short-url/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewShortService(c *gin.Context) ShortUrlSevice {
	return &Service{
		db: ,
	}
}

type Service struct {
	db *gorm.DB
}

func (s *Service) CreateShortUrl(url string, expireAt *time.Time) (string, error) {

	//create sequence_id

	seqID := strconv.FormatInt(int64(seq.ID), 10)
	shortUrl, err := utils.Base62Encode("dog" + seqID)
	if err != nil {
		return "", err
	}

	createTime := time.Now()
	m := models.Url{
		ShortUrl:  shortUrl,
		Url:       url,
		ExpireAt:  expireAt,
		CreatedAt: &createTime,
	}
	if err := s.db.Create(&m).Error; err != nil {
		return "", err
	}
	
	return shortUrl, nil
}

func (s *Service) GetUrl(shortUrl string) (*models.Url, error) {
	u := &models.Url{}
	result := s.db.Where("short_url = ? AND expire_at> ?", shortUrl, time.Now().Format(global.SecLocalTimeFormat)).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
