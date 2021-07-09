package v1

import (
	"errors"
	"net/http"
	"net/url"
	"short-url/cache"
	"short-url/db"
	"short-url/models"
	"short-url/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrParseError = errors.New("parse parameter error")
var ErrExpiredError = errors.New("expired error")
var ErrUrlError = errors.New("parse url error")

func CreateShortUrl(c *gin.Context) {
	req := &models.ShortURLRequest{}

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

	//create sequence_id and encrypt
	seqID, err := db.RetriveDBAccessModel().CreateSequenceAndGetID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	seqIDStr := strconv.FormatInt(int64(seqID), 10)
	shortUrl, err := utils.Base62Encode("shorturl" + seqIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	err = db.RetriveDBAccessModel().CreateUrl(shortUrl, req.Url, req.ExpireAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ShortUrlResponse{ID: shortUrl, Url: c.Request.Host + "/" + shortUrl})
}

func GetOriginalUrl(c *gin.Context) {

	shortUrl := c.Param("shortUrl")

	cacheUrl, found := cache.RetriveCacheAccessModel().Find(shortUrl)
	if found {
		if cacheUrl != "" {
			c.Redirect(http.StatusMovedPermanently, cacheUrl)
		}
	}

	url, err := db.RetriveDBAccessModel().GetUrl(shortUrl)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"err": gorm.ErrRecordNotFound.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	//insert to cache
	err = cache.RetriveCacheAccessModel().Insert(shortUrl, url.Url, url.ExpireAt)
	if err != nil {
		//writing log
	}
	c.Redirect(http.StatusMovedPermanently, url.Url)

}
