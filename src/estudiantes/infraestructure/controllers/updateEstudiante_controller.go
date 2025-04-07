package controllers


import (
	"ZeroDay_Go/src/estudiantes/application"
	"ZeroDay_Go/src/estudiantes/domian/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEstudianteController struct {
	useCase *application.UpdateEstudiante
}

func NewUpdateEstudianteController(useCase *application.UpdateEstudiante) *UpdateEstudianteController {
	return &UpdateEstudianteController{useCase: useCase}
}

func (ctrl *UpdateEstudianteController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var data struct {
		Nombre    string `json:"nombre" binding:"required"`
		Matricula string `json:"matricula" binding:"required"`
		Correo    string `json:"correo" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEstudiante := entities.NewEstudiante(data.Nombre, data.Matricula, data.Correo)

	err = ctrl.useCase.Execute(id, updatedEstudiante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estudiante actualizado con éxito"})
}
