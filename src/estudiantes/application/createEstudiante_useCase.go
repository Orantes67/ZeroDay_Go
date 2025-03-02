package application

import (
	
	"ZeroDay_Go/src/estudiantes/domain"
) 


type CreatEstudiante struct {
	db domain.IEstudiante
}

func NewEstudiante(db domain.IEstudiante) *CreatEstudiante {
	return &CreatEstudiante{db: db}
}

func (cp *CreatEstudiante) Execute() {
	cp.db.Save()
}
