package infraestructure

import (
	
	"ZeroDay_Go/src/maestros/application"
	"ZeroDay_Go/src/maestros/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	ps := NewMySQL()

	createMaestroService := application.NewMaestro(ps)
	listMaestroService := application.NewListMaestro(ps)
	updateMaestroService := application.NewUpdateMaestro(ps)
	deleteMaestroService := application.NewDeleteMaestro(ps)

	createMaestroController := controllers.NewCreateMaestroController(createMaestroService)
	listMaestroController := controllers.NewListMaestrosController(listMaestroService)
	updateMaestroController := controllers.NewUpdateMaestroController(updateMaestroService)
	deleteMaestroController := controllers.NewDeleteMaestroController(deleteMaestroService)

	MaestroRoutes(router, MaestroHandler{
		Create: createMaestroController,
		Get:    listMaestroController,
		Edit:   updateMaestroController,
		Delete: deleteMaestroController,
	})
}