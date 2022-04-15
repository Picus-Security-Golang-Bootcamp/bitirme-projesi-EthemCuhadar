package database

import (
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-EthemCuhadar/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect takes configuration parameters from config file and connect to Postgres
// database. Afterwards, it pings the database whether database reponses for reading
// and writing.
func Connect(cfg *config.Config) *gorm.DB {
	zap.L().Debug("database.bd.Connect")

	// Connection
	db, err := gorm.Open(postgres.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("database.db.Connect - open error", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("database.db.Connect - origin error", zap.Error(err))
	}

	// Set config parameters
	origin.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)

	// Ping
	if err := origin.Ping(); err != nil {
		zap.L().Fatal("database.db.Connect - cannot ping database", zap.Error(err))
	}
	return db
}

// Close terminates the connection of Postgres database.
func Close(db *gorm.DB) {
	zap.L().Debug("database.db.Close")
	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("database.db.Close", zap.Error(err))
	}
	origin.Close()
}
