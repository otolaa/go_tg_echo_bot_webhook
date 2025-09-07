package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	TOKEN string
	URL   string
)

const (
	VERSION string = "0.0.1"
	URL_API string = "https://api.telegram.org/bot"
)

func init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "TOKEN=") {
			TOKEN = strings.ReplaceAll(s, "TOKEN=", "")
		}

		if strings.Contains(s, "URL=") {
			URL = strings.ReplaceAll(s, "URL=", "")
		}
	}

	err = setWebhook()
	if err != nil {
		panic(err)
	}
}

// write to files
func writeJson(data any, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("error marchal %s: %w", filename, err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error write file %s: %w", filename, err)
	}

	fmt.Printf("JSON data write to file ~ %s\n", filename)
	return nil
}
