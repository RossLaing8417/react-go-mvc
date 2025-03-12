package models

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Addresses []Address
type Address struct {
	ID              uint64    `gorm:"not null;primaryKey;unique;autoIncrement"`
	CreatedDateTime time.Time `gorm:"not null;autoCreateTime"`
	UpdatedDateTime time.Time `gorm:"not null;autoUpdateTime"`
	BusinessID      uint64    `gorm:"not null;index:idx1_address"`
	Business        Business  `gorm:"references:ID"`
	StreetNumber    string    `gorm:""`
	Street          string    `gorm:"not null"`
	Town            string    `gorm:"not null;index:idx2_address,priority:2"`
	PostCode        string    `gorm:"not null;index:idx2_address,priority:1"`
}

func (record *Address) sanitize() {
	record.StreetNumber = strings.TrimSpace(record.StreetNumber)
	record.Street = strings.TrimSpace(record.Street)
	record.Town = strings.TrimSpace(record.Town)
	record.PostCode = strings.TrimSpace(record.PostCode)
}

func (record *Address) validate() error {
	if record.BusinessID == 0 {
		return fmt.Errorf("Address must be linked to a business")
	}
	if record.Street == "" {
		return fmt.Errorf("Street is requried")
	}
	if record.Town == "" {
		return fmt.Errorf("Town is requried")
	}
	if record.PostCode == "" {
		return fmt.Errorf("Post code is requried")
	}
	return nil
}

func GetAddressById(db *gorm.DB, id uint64) (Address, error) {
	record := Address{}

	result := db.First(&record, id)
	if result.Error != nil {
		return Address{}, result.Error
	}

	return record, nil
}

func GetAddressesForBusiness(db *gorm.DB, businessId uint64) (Addresses, error) {
	records := Addresses{}

	result := db.Where("business_id = ?", businessId).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (record *Address) Create(db *gorm.DB) error {
	record.sanitize()

	if err := record.validate(); err != nil {
		return err
	}

	result := db.Create(&record)

	return result.Error
}

func (record *Address) Update(db *gorm.DB) error {
	update, err := GetAddressById(db, record.ID)
	if err != nil {
		return err
	}

	update.StreetNumber = record.StreetNumber
	update.Street = record.Street
	update.Town = record.Town
	update.PostCode = record.PostCode

	update.sanitize()
	if err := update.validate(); err != nil {
		return err
	}

	result := db.Where("id = ?", record.ID).Updates(&update)
	if result.Error != nil {
		return result.Error
	}

	*record = update

	return nil
}

func (record *Address) Delete(db *gorm.DB) error {
	del, err := GetAddressById(db, record.ID)
	if err != nil {
		return err
	}

	result := db.Where("id = ?", record.ID).Delete(&del)
	return result.Error
}
