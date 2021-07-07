package initialize

import (
	"fmt"
	"log"
	"tiny-url/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.New()
	err := router.InitRouter(r)
	if err != nil {
		panic(fmt.Errorf("router initialization fail: %w", err))
	}
	log.Printf("router initialization success")
	return r
}
