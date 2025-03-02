package entities

type estudiante struct {
	id int32
	name string
	matricula string
}

func NewEstudiante(name string, matricula string) *estudiante{
	return &estudiante{id: 1, name: name, matricula: matricula}
}

func (e *estudiante) GetName() string {
	return e.name
}

func (e *estudiante) SetName(name string) {
	e.name = name
}