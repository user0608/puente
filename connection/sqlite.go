package connection

import (
	"context"
	"log/slog"
	"path"
	"puente/appconfig"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type StorageManager interface {
	Conn(ctx context.Context) *gorm.DB
	WithTx(ctx context.Context, fc func(ctx context.Context) error) error
}

func NewConnection(c appconfig.AppConfig) (StorageManager, error) {
	var sqliteconn = sqliteConn{databaseDir: c.DataPath(), logLevel: c.DBLogsLevel()}
	if err := sqliteconn.openConnection(); err != nil {
		return nil, err
	}
	return &sqliteconn, nil
}

type sqliteConn struct {
	databaseDir string
	logLevel    string
	conn        *gorm.DB
}

func (*sqliteConn) level(s string) logger.LogLevel {
	switch s {
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	}
	return logger.Silent
}
func (s *sqliteConn) openConnection() error {
	var dataBasePath = path.Join(s.databaseDir, "data.db")
	conn, err := gorm.Open(sqlite.Open(dataBasePath), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(s.level(s.logLevel)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		slog.Error("database error", "error", err, "db_path", dataBasePath)
		return err
	}
	s.conn = conn
	return nil
}

type key int

var context_connextion_key key

func (c *sqliteConn) Conn(ctx context.Context) *gorm.DB {
	value := ctx.Value(context_connextion_key)
	if value == nil {
		return c.conn.WithContext(ctx)
	}
	connection, ok := value.(*gorm.DB)
	if !ok {
		return connection.WithContext(ctx)
	}
	return connection
}

func (c *sqliteConn) WithTx(ctx context.Context, txFn func(ctx context.Context) error) error {
	if ctx.Value(context_connextion_key) != nil {
		return txFn(ctx) // returns the current transaction
	}
	return c.Conn(ctx).Transaction(func(tx *gorm.DB) error {
		return txFn(context.WithValue(ctx, context_connextion_key, tx))
	})
}
