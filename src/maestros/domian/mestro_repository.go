package domian

import (

	"ZeroDay_Go/src/maestros/domian/entities")

type IMaestro interface {
	Save(maestro *entities.Maestro) error
	GetAll() ([]*entities.Maestro, error)
	Update(id int,updateMaestro *entities.Maestro) error
	Delete(id int) error
}