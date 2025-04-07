package controllers

import (
	"ZeroDay_Go/src/materias/application"
	"ZeroDay_Go/src/materias/domian/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMateriaController struct {
	useCase *application.UpdateMateria
}

func NewUpdateMateriaController(useCase *application.UpdateMateria) *UpdateMateriaController {
	return &UpdateMateriaController{useCase: useCase}
}

func (ctrl *UpdateMateriaController) Execute(c *gin.Context) {
	
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var data struct {
		Nombre string `json:"nombre" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Crear instancia de Materia antes de llamar a Execute
	updateMateria := entities.NewMateria(data.Nombre, data.Descripcion)

	err = ctrl.useCase.Execute(id, updateMateria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Materia actualizada con éxito"})

}
