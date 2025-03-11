package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RossLaing8417/react-go-mvc/server/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type businessController struct {
	Db *gorm.DB
}

func NewBusinessController(db *gorm.DB) *businessController {
	return &businessController{
		Db: db,
	}
}

type createBusinessParams struct {
	Name               string `json:"name"`
	VATNumber          uint64 `json:"vat_number"`
	RegistrationNumber string `json:"registration_number"`
}

func (params *createBusinessParams) toModel() models.Business {
	return models.Business{
		Name:               params.Name,
		VATNumber:          params.VATNumber,
		RegistrationNumber: params.RegistrationNumber,
	}
}

type updateBusinessParams struct {
	ID                 uint64 `json:"-"`
	Name               string `json:"name"`
	VATNumber          uint64 `json:"vat_number"`
	RegistrationNumber string `json:"registration_number"`
}

func (params *updateBusinessParams) toModel(id uint64) models.Business {
	return models.Business{
		ID:                 id,
		Name:               params.Name,
		VATNumber:          params.VATNumber,
		RegistrationNumber: params.RegistrationNumber,
	}
}

type businessDTO struct {
	ID                 uint64       `json:"id"`
	CreatedDateTime    time.Time    `json:"created_datetime"`
	UpdateDateTime     time.Time    `json:"updated_datetime"`
	Name               string       `json:"name"`
	VATNumber          uint64       `json:"vat_number"`
	RegistrationNumber string       `json:"registration_number"`
	Addresses          []addressDTO `json:"addresses"`
}

func businessFromModel(record *models.Business) businessDTO {
	addresses := make([]addressDTO, len(record.Addresses))
	for i, address := range record.Addresses {
		addresses[i] = addressFromModel(&address)
	}
	return businessDTO{
		ID:                 record.ID,
		CreatedDateTime:    record.CreatedDateTime,
		UpdateDateTime:     record.UpdatedDateTime,
		Name:               record.Name,
		VATNumber:          record.VATNumber,
		RegistrationNumber: record.RegistrationNumber,
		Addresses:          addresses,
	}
}

func (controller *businessController) CreateBusiness(c *fiber.Ctx) error {
	params := createBusinessParams{}
	if err := c.BodyParser(&params); err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse body")
	}

	record := params.toModel()

	err := record.Create(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to create record")
	}

	return c.Status(http.StatusCreated).JSON(businessFromModel(&record))
}

func (controller *businessController) GetBusinesses(c *fiber.Ctx) error {
	records, err := models.GetBusinesses(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to retrieve records")
	}

	dtos := make([]businessDTO, len(records))
	for i, record := range records {
		dtos[i] = businessFromModel(&record)
	}

	return c.Status(http.StatusOK).JSON(dtos)
}

func (controller *businessController) GetBusiness(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse parameter")
	}

	record, err := models.GetBusinessById(controller.Db, id)
	if err != nil {
		return respondError(c, http.StatusNotFound, err, "Failed to retrieve record")
	}

	return c.Status(http.StatusOK).JSON(businessFromModel(&record))
}

func (controller *businessController) UpdateBusiness(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse parameter")
	}

	params := updateBusinessParams{}
	if err := c.BodyParser(&params); err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse body")
	}

	record := params.toModel(id)

	err = record.Update(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to update record")
	}

	return c.Status(http.StatusOK).JSON(businessFromModel(&record))
}

func (controller *businessController) DeleteBusiness(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse parameter")
	}

	record, err := models.GetBusinessById(controller.Db, id)
	if err != nil {
		return respondError(c, http.StatusNotFound, err, "Failed to retrieve record")
	}

	err = record.Delete(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to delete record")
	}

	return c.SendStatus(http.StatusNoContent)
}
