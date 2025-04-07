package infraestructure

import (

	"ZeroDay_Go/src/maestros/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type MaestroHandler struct{
	Create *controllers.CreateMaestroController
	Get *controllers.ListMaestrosController
	Edit *controllers.UpdateMaestroController
	Delete *controllers.DeleteMaestroController
}
func MaestroRoutes(router *gin.Engine, handler MaestroHandler) {
	maestrosGroup := router.Group("/maestros")
	{
		maestrosGroup.POST("/", handler.Create.Execute)
		maestrosGroup.GET("/", handler.Get.Execute)
		maestrosGroup.PUT("/:id", handler.Edit.Execute)
		maestrosGroup.DELETE("/:id", handler.Delete.Execute)
	}
}