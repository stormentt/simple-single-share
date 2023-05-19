package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getShare(c *gin.Context) {
	id := c.Param("id")

	if len(id) != 36 {
		c.String(http.StatusBadRequest, "invalid share id")
		return
	}

	share, err := shareStore.GetShare(id)
	if err != nil {
		switch err.(type) {
		case *ShareNotFoundError:
			c.String(http.StatusNotFound, "share not found")
		default:
			c.String(http.StatusInternalServerError, "unable to retrieve share due to internal server error. check the server logs")
		}

		log.WithFields(log.Fields{
			"error": err,
		}).Error("unable to retrieve share")
		return
	}

	shareStore.DeleteShare(id)
	c.Data(http.StatusOK, share.ContentType, share.Content)
}
