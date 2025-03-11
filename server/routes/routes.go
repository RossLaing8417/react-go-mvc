package routes

import (
	"github.com/RossLaing8417/react-go-mvc/server/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(api fiber.Router, db *gorm.DB) {

	businessController := controllers.NewBusinessController(db)

	api.Post("/business", businessController.CreateBusiness)
	api.Get("/business", businessController.GetBusinesses)
	api.Get("/business/:id", businessController.GetBusiness)
	api.Put("/business/:id", businessController.UpdateBusiness)
	api.Delete("/business/:id", businessController.DeleteBusiness)

	// Address Routes
	// api.Post("/address")
	// api.Get("/address")
	// api.Patch("/address/:id")
	// api.Delete("/address/:id")
}
