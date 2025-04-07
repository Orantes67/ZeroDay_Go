package polling

import (
	"ZeroDay_Go/src/notificaciones/domain/entities"
	"ZeroDay_Go/src/notificaciones/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	nuevaMateria bool
	materiaMutex sync.Mutex
)

func SetNuevaMateria() {
	materiaMutex.Lock()
	nuevaMateria = true
	materiaMutex.Unlock()
}

func ShortPollingHandler(repo *infrastructure.NotificationRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		materiaMutex.Lock()
		defer materiaMutex.Unlock()

		if nuevaMateria {
			nuevaMateria = false // Reset el estado

			// Guardar en la base de datos
			repo.Save(&entities.Notification{
				Tipo:    "materia",
				Mensaje: "Se ha creado una nueva materia",
			})

			c.JSON(http.StatusOK, gin.H{
				"nuevaMateria": true,
				"mensaje":      "Se ha creado una nueva materia",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"nuevaMateria": false})
	}
}
