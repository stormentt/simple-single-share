package server

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func postShare(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to read request body")
		c.String(http.StatusInternalServerError, "unable to post share due to internal server error. check the server logs")
		return
	}

	uuid, err := shareStore.AddShare(c.ContentType(), data)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to add share")
		c.String(http.StatusInternalServerError, "unable to post share due to internal server error. check the server logs")
		return
	}

	c.String(http.StatusOK, "%s", uuid)
}
