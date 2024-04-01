package migrate

import (
	"context"
	"log/slog"
	"puente/connection"
	"puente/pkg/models"
)

func ExecMigration(manager connection.StorageManager) error {
	var conn = manager.Conn(context.Background())
	if err := conn.AutoMigrate(&models.WorkNode{}); err != nil {
		slog.Error("error while migrating the database", "error", err)
		return err
	}
	return nil
}
