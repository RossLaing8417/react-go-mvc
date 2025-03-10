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

type CreateAddressParams struct {
	BusinessID   uint64 `json:"business_id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

type UpdateAddressParams struct {
	ID           uint64 `json:"-"`
	BusinessID   uint64 `json:"business_id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

type AddressDTO struct {
	ID              uint64    `json:"-"`
	CreatedDateTime time.Time `json:"created_datetime"`
	UpdateDateTime  time.Time `json:"updated_datetime"`
	BusinessID      uint64    `json:"business_id"`
	StreetNumber    string    `json:"street_number"`
	Street          string    `json:"street"`
	Town            string    `json:"town"`
	PostCode        string    `json:"post_code"`
}

func (record *Address) ToDTO() AddressDTO {
	return AddressDTO{
		ID:              record.ID,
		CreatedDateTime: record.CreatedDateTime,
		UpdateDateTime:  record.UpdatedDateTime,
		BusinessID:      record.BusinessID,
		StreetNumber:    record.StreetNumber,
		Street:          record.Street,
		Town:            record.Town,
		PostCode:        record.PostCode,
	}
}
