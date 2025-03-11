package routes

import (
	"github.com/RossLaing8417/react-go-mvc/server/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(api fiber.Router, db *gorm.DB) {

	business_controller := controllers.NewBusinessController(db)

	api.Post("/business", business_controller.CreateBusiness)
	api.Get("/business", business_controller.GetBusinesses)
	api.Get("/business/:id", business_controller.GetBusiness)
	api.Patch("/business/:id", business_controller.UpdateBusiness)
	api.Delete("/business/:id", business_controller.DeleteBusiness)

	// Address Routes
	// api.Post("/address")
	// api.Get("/address")
	// api.Patch("/address/:id")
	// api.Delete("/address/:id")
}
