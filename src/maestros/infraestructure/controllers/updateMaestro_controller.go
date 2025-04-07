package controllers

import (
	"ZeroDay_Go/src/maestros/application"
	"ZeroDay_Go/src/maestros/domian/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMaestroController struct {
	useCase *application.UpdateMaestro
}

func NewUpdateMaestroController(useCase *application.UpdateMaestro) *UpdateMaestroController {
	return &UpdateMaestroController{useCase: useCase}
}
func (ctrl *UpdateMaestroController) Execute(c *gin.Context) {
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

	updatedMaestro := entities.NewMaestro(data.Nombre, data.Matricula, data.Correo)

	err = ctrl.useCase.Execute(id, updatedMaestro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Maestro actualizado con éxito"})
}