package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Businesses []Business
type Business struct {
	ID                 uint64    `gorm:"primaryKey;unique;autoIncrement;<-create"`
	CreatedDateTime    time.Time `gorm:"not null;autoCreateTime;<-create"`
	UpdatedDateTime    time.Time `gorm:"not null;autoUpdateTime"`
	Name               string    `gorm:"not null;uniqueIndex:ak1_business"`
	VATNumber          uint64    `gorm:"index:idx1_business"`
	RegistrationNumber string    `gorm:"index:idx2_business"`
	Addresses          Addresses `gorm:"foreignKey:BusinessID"`
}

func (record *Business) sanitize() {
	record.Name = strings.TrimSpace(record.Name)
	record.RegistrationNumber = strings.TrimSpace(record.RegistrationNumber)
}

func (record *Business) validate() error {
	if record.Name == "" {
		return fmt.Errorf("Name is required")
	}
	return nil
}

func GetBusinessById(db *gorm.DB, id uint64) (Business, error) {
	record := Business{}

	result := db.First(&record, id)
	if result.Error != nil {
		return Business{}, result.Error
	}
	var err error
	record.Addresses, err = GetAddressesForBusiness(db, record.ID)
	if err != nil {
		return Business{}, err
	}

	return record, nil
}

func GetBusinesses(db *gorm.DB) (Businesses, error) {
	records := Businesses{}

	result := db.Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (record *Business) Create(db *gorm.DB) error {
	record.sanitize()

	if err := record.validate(); err != nil {
		return err
	}

	if record.VATNumber != 0 {
		result := db.Where("vat_number = ?", record.VATNumber).Find(&Business{})
		if result.RowsAffected != 0 {
			return fmt.Errorf("Business already exists with the vat number '%v'", record.VATNumber)
		}
	}

	if record.RegistrationNumber != "" {
		result := db.Where("registration_number = ?", record.VATNumber).Find(&Business{})
		if result.RowsAffected != 0 {
			return fmt.Errorf("Business already exists with the registration number '%v'", record.VATNumber)
		}
	}

	result := db.Create(&record)

	return result.Error
}

func (record *Business) Update(db *gorm.DB) error {
	update, err := GetBusinessById(db, record.ID)
	if err != nil {
		return err
	}

	update.Name = record.Name
	update.VATNumber = record.VATNumber
	update.RegistrationNumber = record.RegistrationNumber

	update.sanitize()
	if err = update.validate(); err != nil {
		return err
	}

	if record.VATNumber != 0 {
		result := db.Where("id != ? AND vat_number = ?", record.ID, record.VATNumber).Find(&Business{})
		if result.RowsAffected != 0 {
			return fmt.Errorf("Business already exists with the vat number '%v'", record.VATNumber)
		}
	}

	if record.RegistrationNumber != "" {
		result := db.Where("id != ? AND registration_number = ?", record.ID, record.VATNumber).Find(&Business{})
		if result.RowsAffected != 0 {
			return fmt.Errorf("Business already exists with the registration number '%v'", record.VATNumber)
		}
	}

	result := db.Where("id = ?", record.ID).Updates(&update)
	if result.Error != nil {
		return result.Error
	}

	*record = update

	return nil
}

func (record *Business) Delete(db *gorm.DB) error {
	del, err := GetBusinessById(db, record.ID)
	if err != nil {
		return err
	}

	if len(del.Addresses) != 0 {
		return fmt.Errorf("Cannot delete businesses while an address still exist")
	}

	result := db.Where("id = ?", record.ID).Delete(&del)
	return result.Error
}
