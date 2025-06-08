package rbg_json

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	OutputPath string `rbg_json:"output_path"`
	InputFile  string `rbg_json:"input_file"`
}

var ConfigInstance *Config

func ReadJson() error {
	file, err := os.Open("config.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = CreateJson()
			if err != nil {
				return err
			}
		}
		fmt.Println(err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	bytes, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(bytes, &ConfigInstance)
	if err != nil {
		return err
	}
	return nil
}
