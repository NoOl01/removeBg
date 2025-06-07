package main

import (
	"fmt"
	"os"
	"removeBg/rbg_json"
	"strings"
)

func ReadFile(config *rbg_json.Config) ([]string, error) {
	file, err := os.Open(config.InputFile)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)

	data, err := os.ReadFile(config.InputFile)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
