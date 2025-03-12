package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RossLaing8417/react-go-mvc/server/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type addressController struct {
	Db *gorm.DB
}

func NewAddressController(db *gorm.DB) *addressController {
	return &addressController{
		Db: db,
	}
}

type addressDTO struct {
	ID              uint64    `json:"id"`
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

func (controller *addressController) GetAddress(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse request id parameter")
	}

	record, err := models.GetAddressById(controller.Db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responseError(c, http.StatusNotFound, err, "Failed to retrieve address record")
		}
		return responseError(c, http.StatusInternalServerError, err, "Failed to retrieve address record")
	}

	return c.Status(http.StatusOK).JSON(addressFromModel(&record))
}

func (controller *addressController) GetAddresses(c *fiber.Ctx) error {
	businessId, err := strconv.ParseUint(c.Query("business_id", "0"), 10, 64)
	if err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse request query parameter")
	}

	records, err := models.GetAddressesForBusiness(controller.Db, businessId)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err, "Failed to retrieve address records")
	}

	dtos := make([]addressDTO, len(records))
	for i, record := range records {
		dtos[i] = addressFromModel(&record)
	}

	return c.Status(http.StatusOK).JSON(dtos)
}

type createAddressParams struct {
	BusinessID   uint64 `json:"business_id"`
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

func (params *createAddressParams) toModel() models.Address {
	return models.Address{
		BusinessID:   params.BusinessID,
		StreetNumber: params.StreetNumber,
		Street:       params.Street,
		Town:         params.Town,
		PostCode:     params.PostCode,
	}
}

func (controller *addressController) CreateAddress(c *fiber.Ctx) error {
	params := createAddressParams{}
	if err := c.BodyParser(&params); err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse body")
	}

	record := params.toModel()

	err := record.Create(controller.Db)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err, "Failed to create address record")
	}

	return c.Status(http.StatusCreated).JSON(addressFromModel(&record))
}

type updateAddressParams struct {
	StreetNumber string `json:"street_number"`
	Street       string `json:"street"`
	Town         string `json:"town"`
	PostCode     string `json:"post_code"`
}

func (params *updateAddressParams) toModel(id uint64) models.Address {
	return models.Address{
		ID:           id,
		StreetNumber: params.StreetNumber,
		Street:       params.Street,
		Town:         params.Town,
		PostCode:     params.PostCode,
	}
}

func (controller *addressController) UpdateAddress(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse request id parameter")
	}

	// FIXME: Update all business and address for inline err if
	params := updateAddressParams{}
	if err := c.BodyParser(&params); err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse request body")
	}

	record := params.toModel(id)

	if err := record.Update(controller.Db); err != nil {
		// FIXME: check err not found
		return responseError(c, http.StatusBadRequest, err, "Failed to update address record")
	}

	return c.Status(http.StatusOK).JSON(addressFromModel(&record))
}

func (controller *addressController) DeleteAddress(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return responseError(c, http.StatusBadRequest, err, "Failed to parse request id parameter")
	}

	record := models.Address{
		ID: id,
	}

	err = record.Delete(controller.Db)
	if err != nil {
		return responseError(c, http.StatusInternalServerError, err, "Failed to delete address record")
	}

	return c.SendStatus(http.StatusNoContent)
}
