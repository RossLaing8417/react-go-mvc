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

type updateBusinessParams struct {
	ID                 uint64 `json:"-"`
	Name               string `json:"name"`
	VATNumber          uint64 `json:"vat_number"`
	RegistrationNumber string `json:"registration_number"`
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

	record := models.Business{
		Name:               params.Name,
		VATNumber:          params.VATNumber,
		RegistrationNumber: params.RegistrationNumber,
	}

	err := record.Create(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to create record")
	}

	return c.Status(http.StatusCreated).JSON(businessDTO{
		ID:                 record.ID,
		CreatedDateTime:    record.CreatedDateTime,
		UpdateDateTime:     record.UpdatedDateTime,
		Name:               record.Name,
		VATNumber:          record.VATNumber,
		RegistrationNumber: record.RegistrationNumber,
		Addresses:          []addressDTO{},
	})
}

func (controller *businessController) GetBusinesses(c *fiber.Ctx) error {
	businesses, err := models.FindBusinesses(controller.Db)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to retrieve records")
	}

	// dtos := make([]Business, len(records))
	// for i, record := range records {
	// 	dtos[i] = record.ToDTO()
	// }

	return c.Status(http.StatusOK).JSON(businesses)
}

func (controller *businessController) GetBusiness(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse parameter")
	}

	record, err := models.FindBusiness(controller.Db, id)
	if err != nil {
		return respondError(c, http.StatusNotFound, err, "Failed to retrieve record")
	}

	return c.Status(http.StatusOK).JSON(record)
}

func (controller *businessController) UpdateBusiness(c *fiber.Ctx) error {
	return nil
}

func (controller *businessController) DeleteBusiness(c *fiber.Ctx) error {
	return nil
}
