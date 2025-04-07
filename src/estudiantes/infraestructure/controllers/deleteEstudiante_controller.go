package controllers

import (
	"ZeroDay_Go/src/estudiantes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEstudianteController struct {
	useCase *application.DeleteEstudiante
}

func NewDeleteEstudianteController(useCase *application.DeleteEstudiante) *DeleteEstudianteController {
	return &DeleteEstudianteController{useCase: useCase}
}

func (ctrl *DeleteEstudianteController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado con éxito"})
}
