package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/RossLaing8417/react-go-mvc/server/database"
	"github.com/RossLaing8417/react-go-mvc/server/database/migrations"
	"github.com/RossLaing8417/react-go-mvc/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	flog "github.com/gofiber/fiber/v2/log"
)

func main() {
	configFile := flag.String("config-file", "config.json", "Configuration file")
	flag.Parse()

	opts, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	db, err := database.Connect(opts.DBOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	log.Println("Running auto migration...")
	if err := migrations.AutoMigrate(db); err != nil {
		log.Fatalln(err)
	}
	log.Println("Auto migration complete...")

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: logger.ConfigDefault.Format + " \n>>>>>\n${body}\n-----\n${resBody}\n<<<<<\n",
	}))
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
	LogLevel  string             `json:"log_level"`
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

	switch opts.LogLevel {
	case "Trace":
		flog.SetLevel(flog.LevelTrace)
	case "Debug":
		flog.SetLevel(flog.LevelDebug)
	case "Info":
		flog.SetLevel(flog.LevelInfo)
	case "Warn":
		flog.SetLevel(flog.LevelWarn)
	case "Error":
		flog.SetLevel(flog.LevelError)
	case "Fatal":
		flog.SetLevel(flog.LevelFatal)
	case "Panic":
		flog.SetLevel(flog.LevelPanic)
	case "":
	default:
		return Options{}, fmt.Errorf("Invalid log level %s", opts.LogLevel)
	}

	return opts, nil
}
