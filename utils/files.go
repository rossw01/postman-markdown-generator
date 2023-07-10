package utils

import (
	"fmt"
	"os"
)

func WriteToFile(outputPath string, text string) error {
	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed opening file: '%s'\n", outputPath)
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = fmt.Fprintln(file, text)
	if err != nil {
		return err
	}

	return nil
}

func ValidateOutput(path string) bool {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("Output path '%s' is invalid because that filepath is already in use.\n", os.Args[2])
		return false
	}
	var d []byte
	if err := os.WriteFile(path, d, 0644); err == nil {
		os.Remove(path)
		return true
	}
	fmt.Printf("Output path '%s' is invalid because the filepath is invalid/malformed.\n", os.Args[2])
	return false
}
