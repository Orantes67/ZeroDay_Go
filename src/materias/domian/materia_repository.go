package domian

import (

	"ZeroDay_Go/src/materias/domian/entities")

type IMaestro interface {
	Save(materias *entities.Materia) error
	GetAll() ([]*entities.Materia, error)
	Update(id int,updateMateria *entities.Materia) error
	Delete(id int) error
}