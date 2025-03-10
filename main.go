package main

import (
	"flag"
	"log"

	"github.com/RossLaing8417/react-go-mvc/config"
	"github.com/RossLaing8417/react-go-mvc/server/database"
	"github.com/RossLaing8417/react-go-mvc/server/models"
	"github.com/RossLaing8417/react-go-mvc/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config_file := *flag.String("config-file", "config.json", "Configuration file")
	flag.Parse()

	opts, err := config.ReadConfig(config_file)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	err = database.Connect(opts.DBOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	// TODO: Where to put this...
	database.Instance().AutoMigrate(
		&models.Business{},
		&models.Address{},
	)

	app := fiber.New()
	// app.Use(middleware.ModelValidations())
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(recover.New(recover.ConfigDefault))

	if opts.ApiPrefix == "" {
		routes.Setup(app)
	} else {
		routes.Setup(app.Group(opts.ApiPrefix))
	}

	log.Fatalln(app.Listen(opts.Host + ":" + opts.Port))
}
