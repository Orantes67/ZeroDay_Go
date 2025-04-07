
package infrastructure

import (
	"ZeroDay_Go/src/notificaciones/domain/entities"
	"database/sql"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) Save(notification *entities.Notification) error {
	query := "INSERT INTO notificaciones (tipo, mensaje, fecha) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, notification.Tipo, notification.Mensaje, notification.Fecha)
	return err
}
