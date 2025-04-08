package polling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LongPollingMaestro(c *gin.Context) {
	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Timeout sin nuevos maestros"})
			return
		case <-time.After(1 * time.Second):
			if CheckAndResetNewMaestro() {
				c.JSON(http.StatusOK, gin.H{"mensaje": "Nuevo maestro agregado"})
				return
			}
		}
	}
}
