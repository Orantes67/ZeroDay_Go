package application

import (
	"ZeroDay_Go/src/maestros/domian"
	"ZeroDay_Go/src/maestros/domian/entities"
) 

type CreateMaestro struct {
	db  domian.IMaestro
}

func NewMaestro(db domian.IMaestro) *CreateMaestro {
	return &CreateMaestro{db: db}
}

func (cpay *CreateMaestro) Execute(NewMaestro *entities.Maestro) error {
	err := cpay.db.Save(NewMaestro)
	if err != nil {
		return err
	}
	return nil
}
