package infraestructure

import (
	"ZeroDay_Go/src/estudiantes/application"
	"ZeroDay_Go/src/estudiantes/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine){

	ps := NewMySQL()

	createEstudiantesService := application.NewEstudiante(ps)
	listEstudiantesService := application.NewListEstudiante(ps)
	updateEstudiantesService := application.NewUpdateEstudiante(ps)
	deleteEstudiantesService := application.NewDeleteEstudiante(ps)

	createEstudiantesController := controllers.NewCreateEstudianteController(createEstudiantesService)
	listEstudiantesController := controllers.NewListEstudiantesController(listEstudiantesService)
	updateEstudiantesController := controllers.NewUpdateEstudianteController(updateEstudiantesService)
	deleteEstudiantesServiceController := controllers.NewDeleteEstudianteController(deleteEstudiantesService)

	EstudiantesRoutes(router, EstudianteHandler{
		Create: createEstudiantesController, 
		Get:    listEstudiantesController,  
		Edit:   updateEstudiantesController,
		Delete: deleteEstudiantesServiceController,
	})
}