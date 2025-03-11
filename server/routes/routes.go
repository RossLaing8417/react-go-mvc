package routes

import (
	"github.com/RossLaing8417/react-go-mvc/server/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(api fiber.Router, db *gorm.DB) {

	businessController := controllers.NewBusinessController(db)
	addressController := controllers.NewAddressController(db)

	api.Post("/business", businessController.CreateBusiness)
	api.Get("/business", businessController.GetBusinesses)
	api.Get("/business/:id", businessController.GetBusiness)
	api.Put("/business/:id", businessController.UpdateBusiness)
	api.Delete("/business/:id", businessController.DeleteBusiness)

	api.Post("/address", addressController.CreateAddress)
	api.Get("/address", addressController.GetAddresses)
	api.Get("/address/:id", addressController.GetAddress)
	api.Put("/address/:id", addressController.UpdateAddress)
	api.Delete("/address/:id", addressController.DeleteAddress)
}
