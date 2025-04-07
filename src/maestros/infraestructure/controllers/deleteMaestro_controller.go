package controllers

import (
	"ZeroDay_Go/src/maestros/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type DeleteMaestroController struct {
	useCase *application.DeleteMaestro
}
func NewDeleteMaestroController(useCase *application.DeleteMaestro) *DeleteMaestroController {
	return &DeleteMaestroController{useCase: useCase}
}
func (ctrl *DeleteMaestroController) Execute(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Maestro eliminado con éxito"})
}