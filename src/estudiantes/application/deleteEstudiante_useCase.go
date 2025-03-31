package application

import (
	"ZeroDay_Go/src/estudiantes/domian" 
	
)

type DeleteEstudiante struct {
 db domain.IEstudiante
}

func NewDeleteEstudiante(db domain.IEstudiante) *DeleteEstudiante {
	 return &DeleteEstudiante{
  db: db,
 }
}

func (d *DeleteEstudiante) Execute(id int) error {
	// Lógica para eliminar un estudiante
	err := d.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}