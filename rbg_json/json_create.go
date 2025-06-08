package rbg_json

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateJson() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)

	config := Config{
		OutputPath: "",
		InputFile:  "",
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
