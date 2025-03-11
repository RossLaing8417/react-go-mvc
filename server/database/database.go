package database

import (
	"encoding/json"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBOptions struct {
	LogLevel      logger.LogLevel
	DriverOptions DriverOptions
}

type DriverOptions interface {
	Connection() gorm.Dialector
}

type SQLiteOptions struct {
	Path string `json:"path"`
}

type PostgreSQLOptions struct {
	Name string `json:"name"`
}

// Unmarshal JSON into the correct driver ensuring at least one and only one driver is provided
func (opts *DBOptions) UnmarshalJSON(data []byte) error {
	type jsonOptions struct {
		LogLevel   *string            `json:"log_level"`
		SQLite     *SQLiteOptions     `json:"sqlite"`
		PostgreSQL *PostgreSQLOptions `json:"postgresql"`
	}

	jsonOpts := jsonOptions{}
	err := json.Unmarshal(data, &jsonOpts)
	if err != nil {
		return err
	}

	if jsonOpts.SQLite != nil && jsonOpts.PostgreSQL != nil {
		return fmt.Errorf("More than one database driver provided")
	} else if jsonOpts.SQLite != nil {
		opts.DriverOptions = jsonOpts.SQLite
	} else if jsonOpts.PostgreSQL != nil {
		opts.DriverOptions = jsonOpts.PostgreSQL
	} else {
		fmt.Println("No config")
		return fmt.Errorf("Missing database driver option")
	}

	if jsonOpts.LogLevel != nil {
		switch *jsonOpts.LogLevel {
		case "Silent":
			opts.LogLevel = logger.Silent
		case "Error":
			opts.LogLevel = logger.Error
		case "Warn":
			opts.LogLevel = logger.Warn
		case "Info":
			opts.LogLevel = logger.Info
		default:
			return fmt.Errorf("Invalid log level: %s", *jsonOpts.LogLevel)
		}
	} else {
		opts.LogLevel = logger.Error
	}

	return nil
}

func (opts *SQLiteOptions) Connection() gorm.Dialector {
	return sqlite.Open(opts.Path)
}

func (opts *PostgreSQLOptions) Connection() gorm.Dialector {
	return sqlite.Open(opts.Name)
}

func Connect(opts DBOptions) (*gorm.DB, error) {
	db, err := gorm.Open(opts.DriverOptions.Connection(), &gorm.Config{
		Logger: logger.Default.LogMode(opts.LogLevel),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
