package routes

import (
	"github.com/RossLaing8417/react-go-mvc/server/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(api fiber.Router, db *gorm.DB) {
	// Business Routes
	api.Post("/business", controllers.CreateBusiness)
	api.Get("/business", controllers.GetBusinesses)
	api.Get("/business/:id", controllers.GetBusiness)
	api.Patch("/business/:id", controllers.UpdateBusiness)
	api.Delete("/business/:id", controllers.DeleteBusiness)

	// Address Routes
	// api.Post("/address")
	// api.Get("/address")
	// api.Patch("/address/:id")
	// api.Delete("/address/:id")
}
