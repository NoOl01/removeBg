package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"removeBg/rbg_json"
	"strings"
)

func main() {
	fmt.Println("Читаю конфиг")

	err := rbg_json.ReadJson()
	if err != nil {
		log.Panicf("Ошибка при чтении файла: %s\n", err.Error())
	}

	err = CheckDir()
	if err != nil {
		log.Panicf("Ошибка при проверке/создании папки: %s\n", err.Error())
	}

	switch rbg_json.ConfigInstance.InputFile {
	case "":
		fmt.Println("Удаляю фон из картинок")
		input, err := ReadCmd()
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := exec.Command(
			"python3", "rembg/remove_bg.py",
			fmt.Sprintf("./images/%s", input),
			fmt.Sprintf("%s/%s", rbg_json.ConfigInstance.OutputPath, input),
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Готово")
	default:
		input, err := ReadFile(rbg_json.ConfigInstance)
		if err != nil {
			fmt.Println(err)
			return
		}

		filesLen := len(input)
		fmt.Println("Удаляю фон из картинок")

		for i, file := range input {
			fileName := strings.TrimSpace(file)

			cmd := exec.Command(
				"python3", "rembg/remove_bg.py",
				fmt.Sprintf("./images/%s", fileName),
				fmt.Sprintf("%s/%s", rbg_json.ConfigInstance.OutputPath, fileName),
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				return
			}

			percent := float64(i+1) / float64(filesLen)
			barSize := 50
			filled := int(percent * float64(barSize))
			empty := barSize - filled

			bar := fmt.Sprintf("|%s%s| %.0f%%",
				strings.Repeat("_", filled),
				strings.Repeat(" ", empty),
				percent*100,
			)

			fmt.Printf("\r%s", bar)
		}

		fmt.Println("\nГотово")

	}
}
