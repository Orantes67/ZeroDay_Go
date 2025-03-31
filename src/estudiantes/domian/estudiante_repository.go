package domain

import (
	"ZeroDay_Go/src/estudiantes/domian/entities"
)
type IEstudiante interface {
	Save(estudiante *entities.Estudiante)error
	GetAll()([]entities.Estudiante,error)
	Delete(id int)error
	Update(id int,updatedPago *entities.Estudiante) error
}