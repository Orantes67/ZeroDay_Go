package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortPollingAlumno(c *gin.Context) {
	response := gin.H{
		"nuevo_alumno": CheckNewAlumno(),
	}
	c.JSON(http.StatusOK, response)
}
