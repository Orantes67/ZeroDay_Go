package application


import (
	"ZeroDay_Go/src/maestros/domian" 
)

type DeleteMaestro struct {
	db domian.IMaestro
}

func NewDeleteMaestro(db domian.IMaestro) *DeleteMaestro {
	return &DeleteMaestro{
		db: db,}
}
func (d *DeleteMaestro) Execute(id int) error {
	// Lógica para eliminar un maestro
	err := d.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}