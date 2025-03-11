package controllers

import (
	"time"

	"github.com/RossLaing8417/react-go-mvc/server/models"
)

type createAddressParams struct {
	BusinessID   uint64 `json:"business_id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

type updateAddressParams struct {
	ID           uint64 `json:"-"`
	BusinessID   uint64 `json:"business_id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

type addressDTO struct {
	ID              uint64    `json:"-"`
	CreatedDateTime time.Time `json:"created_datetime"`
	UpdateDateTime  time.Time `json:"updated_datetime"`
	BusinessID      uint64    `json:"business_id"`
	StreetNumber    string    `json:"street_number"`
	Street          string    `json:"street"`
	Town            string    `json:"town"`
	PostCode        string    `json:"post_code"`
}

func addressFromModel(record *models.Address) addressDTO {
	return addressDTO{
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
