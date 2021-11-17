package app

import (
	"errors"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	lock   = &sync.Mutex{}
)

func init() {
	router, _ = generateEngine()
}

func StartApplication() {
	mapURLs()
	router.Run(":8080")
}

func generateEngine() (*gin.Engine, error) {
	if router != nil {
		return nil, errors.New("engine already exist")
	}

	lock.Lock()
	defer lock.Unlock()
	router = gin.Default()
	router.Use(cors.Default())

	return router, nil
}
