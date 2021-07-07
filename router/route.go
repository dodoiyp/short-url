package router

import (
	v1 "tiny-url/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) error {
	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("/urls", v1.CreateTinyURL)
	}
	r.GET(":shortUrl", v1.GetOriginalURL)

	return nil
}
