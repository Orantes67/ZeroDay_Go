package polling

import (
	"ZeroDay_Go/src/notificaciones/domain/entities"
	"ZeroDay_Go/src/notificaciones/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var (
	listMaestrosTriggered  = false
	updateAlumnoTriggered  = false
	mu                     sync.Mutex
)

func SetListadoMaestros() {
	mu.Lock()
	defer mu.Unlock()
	listMaestrosTriggered = true
}

func SetActualizacionAlumno() {
	mu.Lock()
	defer mu.Unlock()
	updateAlumnoTriggered = true
}

func LongPollingHandler(repo *infrastructure.NotificationRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := time.After(30 * time.Second)
		ticker := time.Tick(1 * time.Second)

		for {
			select {
			case <-timeout:
				c.JSON(http.StatusOK, gin.H{"message": "Sin cambios detectados"})
				return
			case <-ticker:
				mu.Lock()
				if listMaestrosTriggered {
					listMaestrosTriggered = false
					repo.Save(&entities.Notification{
						Tipo:    "maestro",
						Mensaje: "Se ha hecho un listado de maestros",
					})
					mu.Unlock()
					c.JSON(http.StatusOK, gin.H{"tipo": "maestro", "message": "Listado de maestros detectado"})
					return
				} else if updateAlumnoTriggered {
					updateAlumnoTriggered = false
					repo.Save(&entities.Notification{
						Tipo:    "alumno",
						Mensaje: "Se ha actualizado un alumno",
					})
					mu.Unlock()
					c.JSON(http.StatusOK, gin.H{"tipo": "alumno", "message": "Actualización de alumno detectada"})
					return
				}
				mu.Unlock()
			}
		}
	}
}
