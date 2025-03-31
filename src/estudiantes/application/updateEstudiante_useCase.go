package application

import (
	"ZeroDay_Go/src/estudiantes/domian"
	"ZeroDay_Go/src/estudiantes/domian/entities"
)

type UpdateEstudiante struct {
	db domain.IEstudiante
}

func NewUpdateEstudiante(db domain.IEstudiante) *UpdateEstudiante {
	return &UpdateEstudiante{
		db: db,
	}
}

func (u *UpdateEstudiante) Execute(id int ,updateEstudiante * entities.Estudiante) error {
	// Lógica para actualizar un estudiante
	err := u.db.Update(id,updateEstudiante)
	if err != nil {
		return err
	}
	return nil
}