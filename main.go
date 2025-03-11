package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

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
	auto_migrate := *flag.Bool("auto-migrate", false, "Run auto migration")
	flag.Parse()

	opts, err := ReadConfig(config_file)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	db, err := database.Connect(opts.DBOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	if auto_migrate {
		log.Println("Running auto migration...")
		err := db.AutoMigrate(
			&models.Business{},
			&models.Address{},
		)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Auto migration complete, now exiting...")
		return
	}

	app := fiber.New()
	// app.Use(middleware.ModelValidations())
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(recover.New(recover.ConfigDefault))

	if opts.ApiPrefix == "" {
		routes.Setup(app, db)
	} else {
		routes.Setup(app.Group(opts.ApiPrefix), db)
	}

	log.Fatalln(app.Listen(opts.Host + ":" + opts.Port))
}

type Options struct {
	Host      string             `json:"host"`
	Port      string             `json:"port"`
	ApiPrefix string             `json:"api_prefix"`
	DBOptions database.DBOptions `json:"database"`
}

func ReadConfig(file_name string) (Options, error) {
	opts := Options{
		Host:      "",
		Port:      "8080",
		ApiPrefix: "/api",
	}

	file, err := os.Open(file_name)
	if err != nil {
		return Options{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&opts)
	if err != nil {
		return Options{}, err
	}

	return opts, nil
}
