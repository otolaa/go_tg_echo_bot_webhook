package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	VERSION string = "0.0.1"
	URL_API string = "https://api.telegram.org/bot"
	SUFFIX  string = "~"
)

var (
	TOKEN      string
	URL        string
	SuffixLine string = strings.Repeat(SUFFIX, 39)
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

	err = setMyCommands()
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

// color: 1 red, 2 green, 3 yello, 4 blue, 5 purple, 6 blue
func p(color int, str ...any) {
	SuffixColor := "\033[3" + strconv.Itoa(color) + "m"
	fmt.Printf("%s%s%s", SuffixColor, fmt.Sprint(str...), "\033[0m\n")
}
