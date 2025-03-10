package controllers

import (
	"net/http"
	"strconv"

	"github.com/RossLaing8417/react-go-mvc/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBusiness(c *fiber.Ctx) error {
	params := models.CreateBusinessParams{}

	if err := c.BodyParser(&params); err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse body")
	}

	business, err := models.CreateBusiness(params)
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to create record")
	}

	return c.Status(http.StatusCreated).JSON(business)
}

func GetBusinesses(c *fiber.Ctx) error {
	businesses, err := models.GetBusinesses()
	if err != nil {
		return respondError(c, http.StatusInternalServerError, err, "Failed to retrieve records")
	}

	return c.Status(http.StatusOK).JSON(businesses)
}

func GetBusiness(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, http.StatusBadRequest, err, "Failed to parse parameter")
	}

	business, err := models.GetBusiness(id)
	if err != nil {
		return respondError(c, http.StatusNotFound, err, "Failed to retrieve record")
	}

	return c.Status(http.StatusOK).JSON(business)
}

func UpdateBusiness(c *fiber.Ctx) error {
	return nil
}

func DeleteBusiness(c *fiber.Ctx) error {
	return nil
}
