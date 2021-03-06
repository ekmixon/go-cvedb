package core

import (
	"github.com/google/wire"
	mdb "go.zenithar.org/pkg/db/adapter/mongodb"

	"go.zenithar.org/cvedb/cli/cvedb/internal/config"
	"go.zenithar.org/cvedb/internal/repositories/pkg/mongodb"
	advisoryv1 "go.zenithar.org/cvedb/internal/services/v1/advisory"
	"go.zenithar.org/cvedb/pkg/pagination"
)

// MongoDBConfig declares a Database configuration provider for Wire
func MongoDBConfig(cfg *config.Configuration) *mdb.Configuration {
	return &mdb.Configuration{
		AutoMigrate:      cfg.DB.AutoMigrate,
		ConnectionString: cfg.DB.Hosts,
		DatabaseName:     cfg.DB.Database,
		Username:         cfg.DB.Username,
		Password:         cfg.DB.Password,
	}
}

var mgoRepositorySet = wire.NewSet(
	MongoDBConfig,
	mongodb.RepositorySet,
)

// -----------------------------------------------------------------------------

// PaginationKey returns the pagination encryption key for cursor pagination.
func PaginationKey(cfg *config.Configuration) (pagination.Key, error) {
	return pagination.KeyFromString(cfg.Server.PaginationKey)
}

// -----------------------------------------------------------------------------

// V1ServiceSet is an object provider for wire
var V1ServiceSet = wire.NewSet(
	mgoRepositorySet,
	PaginationKey,
	advisoryv1.New,
)
