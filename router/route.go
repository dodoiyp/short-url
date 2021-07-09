package router

import (
	"net/http"
	"short-url/api"
	v1 "short-url/api/v1"

	"github.com/gin-gonic/gin"
)

// https://stackoverflow.com/questions/29418478/go-gin-framework-cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func InitRouter(r *gin.Engine) error {
	r.Use(CORSMiddleware())
	apiGroup := r.Group("/api")
	version := apiGroup.Group("/v1")
	{
		version.POST("/urls", v1.CreateShortUrl)
	}
	r.Any("/ping", api.Ping)
	r.GET(":shortUrl", v1.GetOriginalUrl)

	return nil
}
