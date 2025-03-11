package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Businesses []Business
type Business struct {
	ID                 uint64    `gorm:"primaryKey;unique;autoIncrement"`
	CreatedDateTime    time.Time `gorm:"not null;autoCreateTime"`
	UpdatedDateTime    time.Time `gorm:"not null;autoUpdateTime"`
	Name               string    `gorm:"not null;uniqueIndex:ak1_business"`
	VATNumber          uint64    `gorm:"index:idx1_business"`
	RegistrationNumber string    `gorm:"index:idx2_business"`
	Addresses          Addresses `gorm:"foreignKey:BusinessID"`
}

func FindBusiness(db *gorm.DB, id uint64) (Business, error) {
	record := Business{}

	result := db.Find(&record, id)
	if result.Error != nil {
		return Business{}, result.Error
	}
	if result.RowsAffected == 0 {
		return Business{}, fmt.Errorf("Could not find business with the id: %d", id)
	}

	return record, nil
}

func FindBusinesses(db *gorm.DB) (Businesses, error) {
	records := Businesses{}

	result := db.Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (record *Business) Create(db *gorm.DB) error {
	result := db.Create(&record)
	return result.Error
}
