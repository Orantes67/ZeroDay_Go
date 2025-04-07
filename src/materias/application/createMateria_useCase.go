package application

import (
	"ZeroDay_Go/src/materias/domian"
	"ZeroDay_Go/src/materias/domian/entities"
) 

type CreateMateria struct{
	db domian.IMaestro
}

func NewCreateMateria(db domian.IMaestro) *CreateMateria {
	return &CreateMateria{
		db: db,
	}
}

func (c *CreateMateria) Execute( NewMateria*entities.Materia) error {
	err := c.db.Save(NewMateria)
	if err != nil {
		return err
	}
	return nil
}