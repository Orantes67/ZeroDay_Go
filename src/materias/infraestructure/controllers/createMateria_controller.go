package controllers

import (

	"ZeroDay_Go/src/materias/application"
	"ZeroDay_Go/src/materias/domian/entities"
	"net/http"

	"github.com/gin-gonic/gin"

)

type CreateMateriaController struct {
	useCase *application.CreateMateria
}
func NewCreateMateriaController(useCase *application.CreateMateria) *CreateMateriaController {
	return &CreateMateriaController{useCase: useCase}
}

func (cp_c *CreateMateriaController) Execute(c *gin.Context) {
	var data struct {
		Nombre string `json:"nombre" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear instancia de Materia antes de llamar a Execute
	materia := entities.NewMateria(data.Nombre, data.Descripcion)

	err := cp_c.useCase.Execute(materia)// Ahora pasamos un solo argumento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Materia creada con éxito"})
}