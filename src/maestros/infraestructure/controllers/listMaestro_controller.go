package controllers

import (
	"ZeroDay_Go/src/maestros/application"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type ListMaestrosController struct {
	useCase *application.ListMaestro
}
func NewListMaestrosController(useCase *application.ListMaestro) *ListMaestrosController {
	return &ListMaestrosController{useCase: useCase}
}
func (ctrl *ListMaestrosController) Execute(c *gin.Context) {
	maestros, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, maestros)
}