package config

import (
	"encoding/json"
	"os"

	"github.com/RossLaing8417/react-go-mvc/server/database"
)

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
