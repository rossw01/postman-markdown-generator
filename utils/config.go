package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"postmanMarkdownGenerator/models"
)

func LoadConfig(config models.Config) (models.Config, error) {
	configJson, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Error reading config.json")
		return models.Config{}, err
	}
	checkValid := json.Valid(configJson)
	if !checkValid {
		fmt.Printf("config.json is formatted incorrectly and cannot be read.")
		return models.Config{}, err
	}
	errNew := json.Unmarshal(configJson, &config)
	if errNew != nil {
		fmt.Println("Unable to build config from config.json, json is valid but formatted unexpectedly.")
		fmt.Println(errNew.Error())
		return models.Config{}, err
	}
	return config, nil
}
