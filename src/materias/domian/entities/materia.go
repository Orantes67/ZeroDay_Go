package entities

type Materia struct {
	ID int
	IDProfessor int
	IDEstudiante int
	Nombre string
	Descripcion string
}

func NewMateria(id, idProfessor, idEstudiante int, nombre, descripcion string) *Materia {
	return &Materia{
		ID: id,
		IDProfessor: idProfessor,
		IDEstudiante: idEstudiante,
		Nombre: nombre,
		Descripcion: descripcion,
	}
}

func (m *Materia) GetID() int {
	return m.ID
}
func (m *Materia) GetIDProfessor() int {
	return m.IDProfessor
}
func (m *Materia) GetIDEstudiante() int {
	return m.IDEstudiante
}
func (m *Materia) GetNombre() string {
	return m.Nombre
}
func (m *Materia) GetDescripcion() string {
	return m.Descripcion
}
func (m *Materia) SetID(id int) {
	m.ID = id
}
func (m *Materia) SetIDProfessor(idProfessor int) {
	m.IDProfessor = idProfessor
}
func (m *Materia) SetIDEstudiante(idEstudiante int) {
	m.IDEstudiante = idEstudiante
}
func (m *Materia) SetNombre(nombre string) {
	m.Nombre = nombre
}
func (m *Materia) SetDescripcion(descripcion string) {
	m.Descripcion = descripcion
}

