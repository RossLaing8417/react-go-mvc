package database

import (
	"encoding/json"
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DBOptions struct {
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
		SQLite     *SQLiteOptions     `json:"sqlite"`
		PostgreSQL *PostgreSQLOptions `json:"postgresql"`
	}

	json_opts := jsonOptions{}
	err := json.Unmarshal(data, &json_opts)
	if err != nil {
		return err
	}

	if json_opts.SQLite != nil && json_opts.PostgreSQL != nil {
		return fmt.Errorf("More than one database driver provided")
	} else if json_opts.SQLite != nil {
		opts.DriverOptions = json_opts.SQLite
	} else if json_opts.PostgreSQL != nil {
		opts.DriverOptions = json_opts.PostgreSQL
	} else {
		fmt.Println("No config")
		return fmt.Errorf("Missing database driver option")
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
	db, err := gorm.Open(opts.DriverOptions.Connection())
	if err != nil {
		return nil, err
	}
	return db, nil
}
