package controllers

import (
	"ZeroDay_Go/src/materias/application"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type ListMateriasController struct {
	useCase *application.ListMateria
}

func NewListMateriasController(useCase *application.ListMateria) *ListMateriasController {
	return &ListMateriasController{useCase: useCase}
}

func (ctrl *ListMateriasController) Execute(c *gin.Context) {
	materias, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, materias)
}
