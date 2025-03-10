package main

import (
	"flag"
	"log"

	"github.com/RossLaing8417/react-go-mvc/config"
	"github.com/RossLaing8417/react-go-mvc/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config_file := *flag.String("config-file", "config.json", "Configuration file")
	flag.Parse()

	opts, err := config.ReadConfig(config_file)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	db, err := opts.DBOptions.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))

	if opts.ApiPrefix == "" {
		routes.Setup(app, db)
	} else {
		routes.Setup(app.Group(opts.ApiPrefix), db)
	}

	log.Fatalln(app.Listen(opts.Host + ":" + opts.Port))
}
