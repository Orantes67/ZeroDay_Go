package serverzeroday

import (
	"ZeroDay_Go/src/notificaciones/infraestructure"
	"database/sql"
	"github.com/gin-gonic/gin"
	"ZeroDay_Go/src/server_ZeroDay/handler/polling"
)

func Run(db *sql.DB) *gin.Engine {
	r := gin.Default()

	notiRepo := infrastructure.NewNotificationRepository(db)

	r.GET("/polling/eventos", polling.LongPollingHandler(notiRepo))

	return r
}
