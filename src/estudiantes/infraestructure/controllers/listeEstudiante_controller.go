package controllers

import (
	"ZeroDay_Go/src/estudiantes/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListEstudiantesController struct {
	useCase *application.ListEstudiante
}

func NewListEstudiantesController(useCase *application.ListEstudiante) *ListEstudiantesController {
	return &ListEstudiantesController{useCase: useCase}
}

func (ctrl *ListEstudiantesController) Execute(c *gin.Context) {
	estudiantes, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, estudiantes)
}
