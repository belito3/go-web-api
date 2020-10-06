package dbsql

import (
	"context"
	"database/sql"
	"github.com/belito3/go-api-codebase/pkg/logger"
	_ "github.com/lib/pq"
	"time"
)


// Config
type Config struct {
	DriverName		string
	DSN				string
	MaxLifetime		int
	MaxOpenConns	int
	MaxIdleConns	int
}

// NewDB DB
func NewDB(c *Config) (*sql.DB, func(), error) {
	// Opening a driver typically with not attempt to connect to the database
	db, err := sql.Open(c.DriverName, c.DSN)
	if err != nil {
		// this will not be a connection error, but a DSN parse error or
		// another initialization error.
		return nil, nil, err
	}
	cleanFunc := func() {
		err := db.Close()
		if err != nil {
			logger.Errorf(context.Background(), "dbsql db close error: %s", err.Error())
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, cleanFunc, err
	}
	db.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	return db, cleanFunc, nil
}
