package main

import (
	"errors"
	"os"
)

func CheckDir() error {
	dir, err := os.Stat("./images")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = CreateDir()
			if err != nil {
				return err
			}
		}
		return err
	}

	if !dir.IsDir() {
		err = CreateDir()
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateDir() error {
	err := os.Mkdir("./images", 0755)
	if err != nil {
		return err
	}
	return nil
}
