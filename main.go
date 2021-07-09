package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"short-url/cache"
	"short-url/config"
	"short-url/db"
	"short-url/logger"
	"short-url/router"

	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	checkcommit = flag.Bool("version", false, "burry code for check version")

	confInfo     *config.Configuration
	gitcommitnum string
)

func checkComimit() {
	log.Println(gitcommitnum)
}

func Init() error {
	flag.Parse()
	// if there is a needed to check git commit num ... print it out and exit
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}

	// read config and pass variables ...
	var err error
	confInfo, err = config.InitConfig()
	if err != nil {
		return fmt.Errorf("Init config err: %v", err)
	}

	// initialize logger
	if err = logger.InitLog(&confInfo.Logs); err != nil {
		return fmt.Errorf("init logger err: %v", err)
	}

	// initialize  mysql
	if err = db.InitDataBase(&confInfo.Mysql); err != nil {
		return fmt.Errorf("init db err: %v", err)
	}

	// initialize cache
	if err = cache.InitCache(&confInfo.Cache); err != nil {
		return fmt.Errorf("init cache err: %v", err)
	}

	return nil
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err: %v", err)
		}
	}()

	err := Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	route := gin.Default()
	router.InitRouter(route)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%v", confInfo.System.Port),
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Info(fmt.Sprintf("http listen : %v\n", err))
			panic(err)
		}
	}()

	gracefulShutdown()
}

func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	logger.Log.Info("awaiting signal")
	<-done
	logger.Log.Info("exiting")
}
