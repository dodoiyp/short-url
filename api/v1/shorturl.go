package v1

import (
	"errors"
	"net/http"
	"net/url"
	"short-url/models"
	"short-url/pkg/cache_service"
	"short-url/pkg/request"
	"short-url/pkg/service"
	"short-url/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrParseError = errors.New("parse parameter error")
var ErrExpiredError = errors.New("expired error")
var ErrUrlError = errors.New("parse url error")

func CreateShortUrl(c *gin.Context) {
	req := &request.ShortURLRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if req.ExpireAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"err": ErrExpiredError.Error()})
		return
	}

	p, err := url.Parse(req.Url)
	if err != nil || (err == nil && (p.Scheme == "" || p.Host == "")) {
		c.JSON(http.StatusBadRequest, gin.H{"err": ErrUrlError.Error()})
		return
	}

	s := service.NewShortService(c)
	key, err := s.NewKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	shortUrl, err := utils.Base62Encode("dog" + key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	m := models.Url{
		ShortUrl:  shortUrl,
		Url:       req.Url,
		ExpireAt:  *req.ExpireAt,
		CreatedAt: time.Now(),
	}
	err = s.CreateShortUrl(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":  shortUrl,
		"url": c.Request.Host + "/" + shortUrl,
	})
}

func GetOriginalUrl(c *gin.Context) {

	shortUrl := c.Param("shortUrl")
	cs := cache_service.New(c)

	cacheUrl, found := cs.Find(shortUrl)
	if found {
		if cacheUrl != "" {
			c.Redirect(http.StatusMovedPermanently, cacheUrl)
		}
	}

	s := service.NewShortService(c)
	url, err := s.GetUrl(shortUrl)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"err": gorm.ErrRecordNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	//insert to cache
	cs.Insert(shortUrl, &cache_service.Url{Url: url.Url, ExpireAt: url.ExpireAt})
	c.Redirect(http.StatusMovedPermanently, url.Url)

}
