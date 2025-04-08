package server

import (
	"ZeroDay_Go/src/server/handler/polling"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/Longpolling/maestro", polling.LongPollingMaestro)
	r.GET("/shortPolling/alumno", polling.ShortPollingAlumno)
}
