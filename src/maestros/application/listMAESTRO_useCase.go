package application

import (
	"ZeroDay_Go/src/maestros/domian"
	"ZeroDay_Go/src/maestros/domian/entities"
)

type ListMaestro struct {
	db domian.IMaestro
}

func NewListMaestro(db domian.IMaestro) *ListMaestro {
	return &ListMaestro{
		db: db,
	}
}

func (l *ListMaestro) Execute() ([]*entities.Maestro, error) {
	// Lógica para listar maestros
	maestros, err := l.db.GetAll()
	if err != nil {
		return nil, err
	}
	return maestros, nil
}