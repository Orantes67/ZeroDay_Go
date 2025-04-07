package infraestructure

import (
	"ZeroDay_Go/src/core"
	"ZeroDay_Go/src/materias/domian/entities"
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
// Método para guardar una nueva materia
func (mysql *MySQL) Save(materia *entities.Materia) error {
	query := "INSERT INTO Materias (nombre,descripcion) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, materia.GetNombre(),materia.GetDescripcion())
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Materia agregada con éxito: %d", rowsAffected)
	}
	return nil
}

func (m *MySQL) GetAll() ([]*entities.Materia, error) {
	query := "SELECT id, nombre, descripcion FROM Materias"
	rows := m.conn.FetchRows(query)

	var materias []*entities.Materia

	for rows.Next() {
		var materia entities.Materia
		if err := rows.Scan(&materia.ID, &materia.Nombre, &materia.Descripcion); err != nil {
			return nil, err
		}
		materias = append(materias, &materia)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return materias, nil
}
// Método para obtener todas las materias

// Método para actualizar una materia
func (mysql *MySQL) Update(id int,materia *entities.Materia) error {
	query := "UPDATE Materias SET nombre = ?, descripcion = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, materia.GetNombre(), materia.GetDescripcion(), materia.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Materia actualizada con éxito: %d", id)
	}else {
		log.Printf("[MySQL] - No se encontró la materia con ID: %d", id)
	}
	return nil
}

// Método para eliminar una materia
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Materias WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Materia eliminada con éxito: %d", id)
	}
	return nil
}
