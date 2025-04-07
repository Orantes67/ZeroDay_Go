package entities

type Estudiante struct {
	ID int32
	Name string
	Matricula string
	Correo string
}

func NewEstudiante(name string, matricula string, correo string) *Estudiante{
	return &Estudiante{ID: 1, Name: name, Matricula: matricula, Correo: correo}
}

func (e *Estudiante) GetName() string {
	return e.Name
}

func (e *Estudiante) SetName(name string) {
	e.Name = name
}

func (e *Estudiante)GetMatricula()string{
	return e.Matricula
}

func (e *Estudiante)SetMatricula(matricula string){
	e.Matricula = matricula
}

func (e *Estudiante)GetCorreo()string{
	return e.Correo
}
func (e *Estudiante)SetCorreo(correo string){
	e.Correo = correo
}
