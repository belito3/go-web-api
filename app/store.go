package app

import (
	"context"
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/app/dep/dbsql"
	"github.com/belito3/go-api-codebase/app/repository/impl"
	"github.com/belito3/go-api-codebase/app/util"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"go.uber.org/dig"
)

func InitStore(container *dig.Container, conf config.AppConfiguration) (func(), error) {
	// Init dbsql db
	cfg2 := conf.DBSQL
	postgresDB, postgresCall, err := dbsql.NewDB(&dbsql.Config{
		DriverName: cfg2.DriverName,
		DSN: cfg2.DSN(),
		MaxLifetime: cfg2.MaxLifeTime,
		MaxIdleConns: cfg2.MaxIdleConns,
		MaxOpenConns: cfg2.MaxOpenConns})
	if err != nil {
		return nil, err
	}

	_ = container.Provide(func() impl.DBTX {
		return postgresDB
	})

	// TODO: gen unique client id
	ctx := context.Background()
	clientId := util.NewID()

	logger.Infof(ctx,"client id: %v", clientId)
	_ = impl.Inject(container)

	return func() {
		postgresCall()
	}, err
}