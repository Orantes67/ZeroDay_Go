package application


import (
	"ZeroDay_Go/src/materias/domian"
	"ZeroDay_Go/src/materias/domian/entities"
) 


type ListMateria struct{
	db domian.IMaestro
}

func NewListMateria(db domian.IMaestro) *ListMateria {
	return &ListMateria{
		db: db,
	}
}

func (l *ListMateria) Execute() ([]*entities.Materia, error) {
	materias, err := l.db.GetAll()
	if err != nil {
		return nil, err
	}
	return materias, nil
}