package controllers

import (
	"ZeroDay_Go/src/materias/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMateriaController struct {
	useCase *application.DeleteMateria
}

func NewDeleteMateriaController(useCase *application.DeleteMateria) *DeleteMateriaController {
	return &DeleteMateriaController{useCase: useCase}
}

func (ctrl *DeleteMateriaController) Execute(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada con éxito"})
}