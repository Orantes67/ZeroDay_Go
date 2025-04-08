package infraestructure

import (
	
	"ZeroDay_Go/src/maestros/domian/entities"
	"ZeroDay_Go/src/core"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

// Constructor para inicializar la conexión a MySQL
func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Método para guardar un nuevo maestro
func (mysql *MySQL) Save(maestro *entities.Maestro) error {
	query := "INSERT INTO Maestros (nombre, matricula, correo) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, maestro.GetCorreo(), maestro.GetMatricula(), maestro.GetNombre())
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Maestro agregado con éxito: %d", rowsAffected)
	}
	return nil
}
func (m *MySQL) GetAll() ([]*entities.Maestro, error) {
	query := "SELECT idMaestros, nombre, matricula, correo FROM Maestros"
	rows := m.conn.FetchRows(query)

	var maestros []*entities.Maestro

	for rows.Next() {
		var maestro entities.Maestro
		if err := rows.Scan(&maestro.ID, &maestro.Nombre, &maestro.Matricula, &maestro.Correo); err != nil {
			return nil, err
		}
		maestros = append(maestros, &maestro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return maestros, nil
}

// Método para obtener todos los maestros

// Método para actualizar un maestro
func (mysql *MySQL) Update(id int ,maestro *entities.Maestro) error {


	query := "UPDATE Maestros SET nombre = ?, matricula = ?, correo = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, maestro.GetNombre(), maestro.GetMatricula(), maestro.GetCorreo(), maestro.GetID())
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Maestro actualizado con éxito: %d", id)
	}else{
		log.Printf("[MySQL] - No se encontró el maestro con ID: %d", id)
	}
	return nil
}
	

// Método para eliminar un maestro
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Maestros WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Maestro eliminado con éxito: %d", id)

	}else{
		log.Printf("[MySQL] - No se encontró el maestro con ID: %d", id)
	}
	return nil
}