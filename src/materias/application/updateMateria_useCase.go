package application

import (
	"ZeroDay_Go/src/materias/domian"
	"ZeroDay_Go/src/materias/domian/entities"
) 


type UpdateMateria struct{
	db domian.IMaestro
}

func NewUpdateMateria(db domian.IMaestro) *UpdateMateria {
	return &UpdateMateria{
		db: db,
	}
}
func (u *UpdateMateria) Execute(id int,updateMateria *entities.Materia) error {
	err := u.db.Update(id, updateMateria)
	if err != nil {
		return err
	}
	return nil
}