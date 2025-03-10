package models

import (
	"fmt"
	"time"

	"github.com/RossLaing8417/react-go-mvc/server/database"
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

type CreateBusinessParams struct {
	Name               string `json:"name"`
	VATNumber          uint64 `json:"vat_number"`
	RegistrationNumber string `json:"registration_number"`
}

type UpdateBusinessParams struct {
	ID                 uint64 `json:"-"`
	Name               string `json:"name"`
	VATNumber          uint64 `json:"vat_number"`
	RegistrationNumber string `json:"registration_number"`
}

type BusinessDTO struct {
	ID                 uint64       `json:"id"`
	CreatedDateTime    time.Time    `json:"created_datetime"`
	UpdateDateTime     time.Time    `json:"updated_datetime"`
	Name               string       `json:"name"`
	VATNumber          uint64       `json:"vat_number"`
	RegistrationNumber string       `json:"registration_number"`
	Addresses          []AddressDTO `json:"addresses"`
}

func CreateBusiness(params CreateBusinessParams) (BusinessDTO, error) {
	record := Business{
		Name:               params.Name,
		VATNumber:          params.VATNumber,
		RegistrationNumber: params.RegistrationNumber,
	}

	result := database.Instance().Create(&record)
	if result.Error != nil {
		return BusinessDTO{}, result.Error
	}

	return BusinessDTO{
		ID:                 record.ID,
		CreatedDateTime:    record.CreatedDateTime,
		UpdateDateTime:     record.UpdatedDateTime,
		Name:               record.Name,
		VATNumber:          record.VATNumber,
		RegistrationNumber: record.RegistrationNumber,
		Addresses:          []AddressDTO{},
	}, nil
}

func GetBusinesses() ([]BusinessDTO, error) {
	records := Businesses{}

	result := database.Instance().Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	dtos := make([]BusinessDTO, len(records))
	for i, record := range records {
		dtos[i] = record.ToDTO()
	}

	return dtos, nil
}

func GetBusiness(id uint64) (BusinessDTO, error) {
	record := Business{}

	result := database.Instance().Find(&record, id)
	if result.Error != nil {
		return BusinessDTO{}, result.Error
	}
	if result.RowsAffected == 0 {
		return BusinessDTO{}, fmt.Errorf("Could not find business with the id: %d", id)
	}

	return record.ToDTO(), nil
}

func (record *Business) ToDTO() BusinessDTO {
	addresses := make([]AddressDTO, len(record.Addresses))
	for i, address := range record.Addresses {
		addresses[i] = address.ToDTO()
	}
	return BusinessDTO{
		ID:                 record.ID,
		CreatedDateTime:    record.CreatedDateTime,
		UpdateDateTime:     record.UpdatedDateTime,
		Name:               record.Name,
		VATNumber:          record.VATNumber,
		RegistrationNumber: record.RegistrationNumber,
		Addresses:          addresses,
	}
}
