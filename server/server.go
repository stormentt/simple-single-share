package server

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var shareStore *ShareStore

func ginLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	c.Next()

	elapsed := time.Since(start)

	log.WithFields(log.Fields{
		"method":  c.Request.Method,
		"path":    path,
		"elapsed": elapsed,
		"status":  c.Writer.Status(),
	}).Info("")
}

func Serve() error {
	r := gin.New()
	r.Use(ginLogger)
	r.Use(gin.Recovery())
	//r.SetTrustedProxies(nil)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.WithFields(log.Fields{
			"method":  httpMethod,
			"handler": handlerName,
		}).Debug(absolutePath)
	}

	r.GET("/share/:id", getShare)
	authRoutes := r.Group("/", CheckAPIKey)
	authRoutes.POST("/share", postShare)

	shareStore = NewShareStore()
	r.Run()

	return nil
}
