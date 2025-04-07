package infraestructure


import (
	"github.com/gin-gonic/gin"
	"ZeroDay_Go/src/estudiantes/infraestructure/controllers"
)


type EstudianteHandler struct{
	Create *controllers.CreateEstudianteController
	Get *controllers.ListEstudiantesController
	Edit *controllers.UpdateEstudianteController
	Delete *controllers.DeleteEstudianteController
}

func EstudiantesRoutes(router *gin.Engine, handler EstudianteHandler) {
	estudiantesGroup := router.Group("/estudiantes")
	{
		estudiantesGroup.POST("/", handler.Create.Execute)
		estudiantesGroup.GET("/", handler.Get.Execute)
		estudiantesGroup.PUT("/:id", handler.Edit.Execute)
		estudiantesGroup.DELETE("/:id", handler.Delete.Execute)
	}
}
