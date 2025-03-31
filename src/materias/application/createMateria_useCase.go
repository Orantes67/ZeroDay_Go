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

func (c *CreateMateria) Execute(id, idProfessor, idEstudiante int, nombre, descripcion string) (*entities.Materia, error) {
	materia := entities.NewMateria(id, idProfessor, idEstudiante, nombre, descripcion)
	err := c.db.Save(materia)
	if err != nil {
		return nil, err
	}
	return materia, nil
}