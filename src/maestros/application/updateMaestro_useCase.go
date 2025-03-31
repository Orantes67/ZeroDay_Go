package application

import (
	"ZeroDay_Go/src/maestros/domian"
	"ZeroDay_Go/src/maestros/domian/entities"
)

type UpdateMaestro struct {
	db domian.IMaestro
}
func NewUpdateMaestro(db domian.IMaestro) *UpdateMaestro {
	return &UpdateMaestro{
		db: db,
	}
}

func (u *UpdateMaestro) Execute(id int,updateMaestro *entities.Maestro) error {
	// Lógica para actualizar un maestro
	err := u.db.Update(id,updateMaestro)
	if err != nil {
		return err
	}
	return nil
}
