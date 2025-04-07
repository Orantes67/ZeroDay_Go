package controllers

import (
	"ZeroDay_Go/src/estudiantes/application"
	"ZeroDay_Go/src/estudiantes/domian/entities"
	"net/http"
	"ZeroDay_Go/src/server/handler/polling"
	"github.com/gin-gonic/gin"
)

type CreateEstudianteController struct {
	useCase *application.CreatEstudiante
}

func NewCreateEstudianteController(useCase *application.CreatEstudiante) *CreateEstudianteController {
	return &CreateEstudianteController{useCase: useCase}
}

func (cp_c *CreateEstudianteController) Execute(c *gin.Context) {
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
	estudiante := entities.NewEstudiante(data.Nombre, data.Matricula, data.Correo)

	err := cp_c.useCase.Execute(estudiante) // Ahora pasamos un solo argumento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	polling.SetNewAlumnoAdded()
	

	c.JSON(http.StatusCreated, gin.H{"message": "Estudiante creado con éxito"})
}