package repository

import (
	"fmt"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectionConf represents an Postgres connection settings
type ConnectionConf struct {
	Username, Password, DBName, DBHost string
	DBPort                             int
}

// ConnectionDB new instance of connection database
type ConnectionDB struct {
	DB *gorm.DB
}

// NewPostgreDB creates a new instance of postgres service
func NewPostgresDB(cfg *ConnectionConf, gormConfig *gorm.Config) (*ConnectionDB, error) {
	if cfg == nil || reflect.DeepEqual(cfg, &ConnectionConf{}) {
		return nil, ErrEmptyConfig
	}
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DBHost, cfg.Username, cfg.Password,
		cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURI,
		PreferSimpleProtocol: true,
	}), gormConfig)
	if err != nil {
		return nil, err
	}
	p := &ConnectionDB{
		DB: db,
	}

	return p, nil
}
