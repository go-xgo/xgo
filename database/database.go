package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	defaultDB *gorm.DB
)

type Config struct {
	Dsn string `json:"dsn"`
}

func Init(cfg *Config) error {
	db, err := New(cfg)
	if err != nil {
		return err
	}
	SetDefaultDB(db)
	return nil
}

func New(cfg *Config) (*gorm.DB, error) {

	if cfg.Dsn == "" {
		return nil, errors.New("empty dsn")
	}
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DefaultDB returns the default db.
func DefaultDB() *gorm.DB {
	return defaultDB
}

// SetDefaultDB sets the default db for package database.
func SetDefaultDB(db *gorm.DB) {
	defaultDB = db
}
