package infraestructure

import (
	"ZeroDay_Go/src/materias/application"
	"ZeroDay_Go/src/materias/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	ps := NewMySQL()

	createMateriaService := application.NewCreateMateria(ps)
	listMateriaService := application.NewListMateria(ps)
	updateMateriaService := application.NewUpdateMateria(ps)
	deleteMateriaService := application.NewDeleteMateria(ps)

	createMateriaController := controllers.NewCreateMateriaController(createMateriaService)
	listMateriaController := controllers.NewListMateriasController(listMateriaService)
	updateMateriaController := controllers.NewUpdateMateriaController(updateMateriaService)
	deleteMateriaController := controllers.NewDeleteMateriaController(deleteMateriaService)

	MateriaRoutes(router, MateriaHandler{
		Create: createMateriaController,
		Get:    listMateriaController,
		Edit:   updateMateriaController,
		Delete: deleteMateriaController,
	})
}
