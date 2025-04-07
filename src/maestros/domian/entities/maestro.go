package entities

type Maestro struct {
	ID        int
	Nombre    string
	Matricula  string
	Correo   string
}

func NewMaestro(nombre string, matricula string, correo string) *Maestro{
	return &Maestro{
		ID:        1,
		Nombre:    nombre,
		Matricula:  matricula,
		Correo:   correo,
	}
}

func (m *Maestro) GetID() int {
	return m.ID
}
func (m *Maestro) GetNombre() string {
	return m.Nombre
}
func (m *Maestro) GetMatricula() string {
	return m.Matricula
}
func (m *Maestro) GetCorreo() string {
	return m.Correo
}
func (m *Maestro) SetID(id int) {
	m.ID = id
}
func (m *Maestro) SetNombre(nombre string) {
	m.Nombre = nombre
}
func (m *Maestro) SetMatricula(matricula string) {
	m.Matricula = matricula
}
func (m *Maestro) SetCorreo(correo string) {
	m.Correo = correo
}
