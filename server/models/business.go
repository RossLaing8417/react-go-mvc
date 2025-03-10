package models

import "time"

type Businesses []Business
type Business struct {
	ID                 uint      `gorm:"primaryKey;unique;autoIncrement"`
	Name               string    `gorm:"not null;uniqueIndex:ak1_business"`
	VATNumber          uint64    `gorm:"index:idx1_business"`
	RegistrationNumber string    `gorm:"index:idx2_business"`
	CreatedTime        time.Time `gorm:"not null;autoCreateTime"`
	UpdatedTime        time.Time `gorm:"not null;autoUpdateTime"`
	Addresses          Addresses `gorm:"foreignKey:BusinessID"`
}
