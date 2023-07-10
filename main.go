package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"postmanMarkdownGenerator/models"
	"postmanMarkdownGenerator/utils"
	"strings"
)

func parseData(inputPath string, outputPath string) {
	jsonData, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Unable to load file '%s'\n", filepath.Base(inputPath))
		return
	}
	checkValid := json.Valid(jsonData)
	if !checkValid {
		fmt.Println("Postman outputted invalid json data. This is unusual and something pretty bad has happened")
		return
	}

	var documentation models.Documentation
	errNew := json.Unmarshal(jsonData, &documentation)
	if errNew != nil {
		fmt.Println("Unable to parse inputted file, file is malformed:")
		fmt.Println(errNew.Error())
		return
	}
	// fmt.Printf("%#v\n", documentation)
	buildMarkdownFile(documentation, outputPath)
}

func buildMarkdownFile(documentation models.Documentation, outputPath string) {
	var config models.Config
	var err error
	utils.InitialiseConfig()
	config, err = utils.LoadConfig(config)
	if err != nil {
		return
	}

	var outputCollectionContents = func(collection models.Collection) {
		for _, endpoint := range collection.Item {
			// INFO
			utils.WriteToFile(outputPath, strings.Replace(config.EndpointName, "{}", endpoint.Name, 1))                       // Name
			utils.WriteToFile(outputPath, strings.Replace(config.EndpointDescription, "{}", endpoint.Request.Description, 1)) // Description
			if !config.UseMethodImage {
				utils.WriteToFile(outputPath, strings.Replace(config.EndpointMethod, "{}", endpoint.Request.Method, 1)) // Request Type
			} else {
				switch endpoint.Request.Method {
				case "PUT":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![PUT Image](%s)", config.PutImageUrl))
					break
				case "POST":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![POST Image](%s)", config.PostImageUrl))
					break
				case "DELETE":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![DELETE Image](%s)", config.DeleteImageUrl))
					break
				case "GET":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![GET Image](%s)", config.GetImageUrl))
					break
				case "OPTIONS":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![OPTIONS Image](%s)", config.OptionsImageUrl))
					break
				case "HEAD":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![HEAD Image](%s)", config.HeadImageUrl))
					break
				case "PATCH":
					utils.WriteToFile(outputPath, fmt.Sprintf("\n![PATCH Image](%s)", config.PatchImageUrl))
					break
				}
			}
			utils.WriteToFile(outputPath, strings.Replace(config.EndpointUrl, "{}", endpoint.Request.Url.Raw, 1)) // Endpoint URL
			// BODY
			if len(endpoint.Request.Body.Raw) > 0 {
				utils.WriteToFile(outputPath, strings.Replace(config.EndpointBodyFormat, "{}", endpoint.Request.Body.Options.Raw.Language, 1))
				utils.WriteToFile(outputPath, strings.Replace(strings.Replace(config.EndpointBody, "{}", endpoint.Request.Body.Options.Raw.Language, 1), "{}", endpoint.Request.Body.Raw, 1))
			}
			utils.WriteToFile(outputPath, "\n<br >\n")
		}
	}

	for _, collection := range documentation.Item {
		err := utils.WriteToFile(outputPath, "# Collection: "+collection.Name)
		if err != nil {
			return
		}
		outputCollectionContents(collection)
	}
	utils.WriteToFile(outputPath, "[Created using PostmanMarkdownGenerator by Ross W.](https://github.com/rossw01/postman-markdown-generator)")
	fmt.Printf("Successfully Generated '%s' at file location: %s \n", filepath.Base(outputPath), outputPath)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(
			"Please include the path to a .json file as a parameter when running this program.\n" +
				"E.g. postmg ./exampleFile.json",
		)
		return
	}

	outputPath := "./" + filepath.Base(os.Args[1])
	lastIndex := strings.LastIndex(outputPath, ".json")
	if lastIndex == -1 {
		fmt.Println("Inputted file doesn't contain a '.json' extension. Please select a different file and try again")
		return
	}
	outputPath = outputPath[:lastIndex] + ".md"

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Printf("WriteToFile with path '%s' doesn't exist. Please try again with an amended filepath.\n", os.Args[1])
		return
	} else if len(os.Args) > 2 {
		if !utils.ValidateOutput(os.Args[2]) {
			return
		} else {
			outputPath = os.Args[2]
		}
	}
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("Output file already exists, try specifying a new output path as a second parameter or deleting the pre-existing file '%s'.\n", filepath.Base(outputPath))
		return
	}
	parseData(os.Args[1], outputPath)
}
