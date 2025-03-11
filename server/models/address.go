package models

import "time"

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
