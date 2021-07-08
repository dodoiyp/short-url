package router

import (
	"short-url/api"
	v1 "short-url/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) error {
	r.Any("/ping", api.Ping)
	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("/urls", v1.CreateShortUrl)
	}
	r.GET(":shortUrl", v1.GetOriginalUrl)

	return nil
}
