package application

import (
	"ZeroDay_Go/src/materias/domian"
)

type DeleteMateria struct {
	db domian.IMaestro
}

func NewDeleteMateria(db domian.IMaestro) *DeleteMateria {
	return &DeleteMateria{
		db: db,
	}
}

func (d *DeleteMateria) Execute(id int) error {
	err := d.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}