package application

import (
	"ZeroDay_Go/src/estudiantes/domian"
	"ZeroDay_Go/src/estudiantes/domian/entities"
)

type ListEstudiante struct {
	db domain.IEstudiante
}

func NewListEstudiante(db domain.IEstudiante) *ListEstudiante {
	return &ListEstudiante{
		db: db,
	}
}

func (l *ListEstudiante) Execute() ([]entities.Estudiante, error) {
	// Lógica para listar estudiantes
	estudiantes, err := l.db.GetAll()
	if err != nil {
		return nil, err
	}
	return estudiantes, nil
}