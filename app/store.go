package app

import (
	"context"
	"database/sql"
	"github.com/belito3/go-web-api/app/config"
	"github.com/belito3/go-web-api/app/repository/impl"
	"github.com/belito3/go-web-api/app/util"
	"github.com/belito3/go-web-api/pkg/logger"
	"go.uber.org/dig"
)

func InitStore(container *dig.Container, conf config.AppConfiguration) (func(), error) {
	// Init dbsql db
	cfg2 := conf.DBSQL
	sqlDB, sqlDBCall, err := impl.NewDB(&impl.Config{
		DriverName: cfg2.DriverName,
		DSN: cfg2.DSN(),
		MaxLifetime: cfg2.MaxLifeTime,
		MaxIdleConns: cfg2.MaxIdleConns,
		MaxOpenConns: cfg2.MaxOpenConns})
	if err != nil {
		return nil, err
	}

	_ = container.Provide(func() *sql.DB {
		return sqlDB
	})

	// TODO: gen unique client id
	ctx := context.Background()
	clientId := util.NewID()

	logger.Infof(ctx,"client id: %v", clientId)
	_ = impl.Inject(container)

	return func() {
		sqlDBCall()
	}, err
}