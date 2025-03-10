package database

import (
	"github.com/RossLaing8417/react-go-mvc/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Instance() *gorm.DB {
	return db
}

func Connect(opts config.DBOptions) error {
	var err error
	db, err = gorm.Open(opts.DriverOptions.Connection())
	if err != nil {
		return err
	}
	return nil
}
