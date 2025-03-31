package entities

type Estudiante struct {
	id int32
	name string
	matricula string
}

func NewEstudiante(name string, matricula string) *Estudiante{
	return &Estudiante{id: 1, name: name, matricula: matricula}
}

func (e *Estudiante) GetName() string {
	return e.name
}

func (e *Estudiante) SetName(name string) {
	e.name = name
}

func (e *Estudiante)GetMatricula()string{
	return e.matricula
}

func (e *Estudiante)SetMatricula(matricula string){
	e.matricula = matricula
}