package infraestructure

import (
	"ZeroDay_Go/src/materias/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type MateriaHandler struct {
	Create *controllers.CreateMateriaController
	Get    *controllers.ListMateriasController
	Edit   *controllers.UpdateMateriaController
	Delete *controllers.DeleteMateriaController
}

func MateriaRoutes(router *gin.Engine, handler MateriaHandler) {
	materiasGroup := router.Group("/materias")
	{
		materiasGroup.POST("/", handler.Create.Execute)
		materiasGroup.GET("/", handler.Get.Execute)
		materiasGroup.PUT("/:id", handler.Edit.Execute)
		materiasGroup.DELETE("/:id", handler.Delete.Execute)
	}
}