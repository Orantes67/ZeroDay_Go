package controllers

import (
	"ZeroDay_Go/src/maestros/application"
	"ZeroDay_Go/src/maestros/domian/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateMaestroController struct {
	useCase *application.CreateMaestro
}

func NewCreateMaestroController(useCase *application.CreateMaestro) *CreateMaestroController {
	return &CreateMaestroController{useCase: useCase}
}

func (cp_c *CreateMaestroController) Execute(c *gin.Context) {
	var data struct {
		Nombre    string `json:"nombre" binding:"required"`
		Matricula string `json:"matricula" binding:"required"`
		Correo    string `json:"correo" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear instancia de Estudiante antes de llamar a Execute
	maestro := entities.NewMaestro(data.Correo,data.Matricula,data.Correo)

	err := cp_c.useCase.Execute(maestro) // Ahora pasamos un solo argumento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Maestro creado con éxito"})
}
