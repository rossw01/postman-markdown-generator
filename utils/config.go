package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"postmanMarkdownGenerator/models"
)

func LoadConfig(config models.Config) (models.Config, error) {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".pmdg", "config.json")
	configJson, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config.json")
		return models.Config{}, errors.New("error reading config.json")
	}
	checkValid := json.Valid(configJson)
	if !checkValid {
		fmt.Printf("config.json is formatted incorrectly and cannot be read.")
		return models.Config{}, errors.New("config.json is formatted incorrectly and cannot be read")

	}
	errNew := json.Unmarshal(configJson, &config)
	if errNew != nil {
		fmt.Println("Unable to build config from config.json, json is valid but formatted unexpectedly.")
		return models.Config{}, errors.New("unable to build config from config.json, json is valid but formatted unexpectedly")
	}
	return config, nil
}

func InitialiseConfig() {
	initialSettings := models.Config{
		EndpointName:        "## End-point: {}",
		EndpointDescription: "{}",
		EndpointMethod:      "#### Method: {}",
		EndpointUrl:         ">```\n>{}\n>```",
		EndpointBodyFormat:  "#### Body(**{}**)",
		EndpointBody:        "```{}\n{}\n```",
		UseMethodImage:      true,
		GetImageUrl:         "https://i.imgur.com/EM0n9ms.png",
		PutImageUrl:         "https://i.imgur.com/c32fs7n.png",
		DeleteImageUrl:      "https://i.imgur.com/kpUiCCt.png",
		PostImageUrl:        "https://i.imgur.com/RSexnGX.png",
		OptionsImageUrl:     "https://i.imgur.com/Oz7Fxtv.png",
		PatchImageUrl:       "https://i.imgur.com/ptO2g5s.png",
		HeadImageUrl:        "https://i.imgur.com/GbCxt6z.png",
	}

	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".pmdg")

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		_ = os.Mkdir(configPath, 0700)
	}

	configPath = filepath.Join(configPath, "config.json")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		file, _ := os.Create(configPath)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Error creating config file")
				return
			}
		}(file)

		settings, _ := json.MarshalIndent(initialSettings, "", "")
		_, err := file.Write(settings)
		if err != nil {
			return
		}
	}
}
