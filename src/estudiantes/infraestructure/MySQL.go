
package infraestructure

import (
	"ZeroDay_Go/src/core"
    "ZeroDay_Go/src/estudiantes/domian/entities"
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

// Método para guardar un nuevo estudiante
func (mysql *MySQL) Save(estudiante *entities.Estudiante) error {
    query := "INSERT INTO Estudiantes (nombre, matricula, correo) VALUES (?, ?, ?)"
    result, err := mysql.conn.ExecutePreparedQuery(query,estudiante.GetCorreo(), estudiante.GetMatricula(), estudiante.GetName() )
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Estudiante agregado con éxito: %d", rowsAffected)
    }
    return nil
}

// Método para obtener todos los estudiantes
func (mysql *MySQL) GetAll() ([]entities.Estudiante, error) {
	query := "SELECT id, nombre, matricula, correo FROM Estudiantes"
	rows := mysql.conn.FetchRows(query)
	
	var estudiantes []entities.Estudiante

	for rows.Next() {
		var estudiante entities.Estudiante
		if err := rows.Scan(&estudiante.ID, &estudiante.Name, &estudiante.Matricula, &estudiante.Correo); err != nil {
			return nil, err
		}
		estudiantes = append(estudiantes, estudiante)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return estudiantes, nil
}

// Método para actualizar un estudiante
func (mysql *MySQL) Update(id int,updateEstudiante *entities.Estudiante) error {
    query := "UPDATE Estudiantes SET nombre = ?, matricula = ?, correo = ? WHERE idEstudiantes = ?"
    result, err := mysql.conn.ExecutePreparedQuery(query, updateEstudiante.GetName(), updateEstudiante.GetMatricula(), updateEstudiante.GetCorreo(), id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Estudiante con ID %d actualizado correctamente", id)
    } else {
        log.Printf("[MySQL] - No se encontró el estudiante con ID %d", id)
    }
    return nil
}

func (mysql *MySQL) Delete(id int) error {
    query := "DELETE FROM Estudiantes WHERE idEstudiantes = ?"
    result, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Estudiante con ID %d eliminado correctamente", id)
    } else {
        log.Printf("[MySQL] - No se encontró el estudiante con ID %d", id)
    }
    return nil
}
