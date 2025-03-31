package application

import (
	 "ZeroDay_Go/src/estudiantes/domian"
	 "ZeroDay_Go/src/estudiantes/domian/entities"
) 


type CreatEstudiante struct {
	db  domain.IEstudiante
}

func NewEstudiante(db domain.IEstudiante) *CreatEstudiante {
	return &CreatEstudiante{db: db}
}

func (cpay *CreatEstudiante)Execute(NewEstudiante *entities.Estudiante) error{
	err := cpay.db.Save(NewEstudiante)
	if err != nil {
		return err
	}
	return nil
}