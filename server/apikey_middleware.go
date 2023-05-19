package server

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CheckAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("API-Key")
	if len(apiKey) == 0 {
		c.String(http.StatusUnauthorized, "No API-Key header provided")
		c.Abort()
		return
	}

	expectedKey := viper.GetString("uploads.api_key")

	log.WithFields(log.Fields{
		"provided": apiKey,
		"expected": expectedKey,
	}).Trace("api key compare")

	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedKey)) != 1 {
		c.String(http.StatusUnauthorized, "Invalid API Key provided")
		c.Abort()
		return
	}

	c.Next()
}
