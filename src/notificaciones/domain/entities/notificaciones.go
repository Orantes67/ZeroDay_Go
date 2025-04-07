
package entities

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Tipo      string    `json:"tipo"` // "maestro" o "alumno"
	Mensaje   string    `json:"mensaje"`
	Fecha     time.Time `json:"fecha"`
}
