package main

import (
	"fmt"
	"os/exec"
	"removeBg/rbg_json"
	"runtime"
	"strings"
)

func main() {
	uos := runtime.GOOS
	var dist string
	if uos == "linux" {
		PrintStatus("Проверяю дистрибутив", "...")
		dist = getLinuxDistro()
		PrintStatus("\rПроверяю дистрибутив", "Done\n")
	}

	PrintStatus("Проверяю установленны ли необходимые файлы", "...")

	err := CheckDependencies(uos, dist)
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintStatus("Читаю конфиг", "...")

	err = rbg_json.ReadJson()
	if err != nil {
		PrintStatus("Читаю конфиг", "Error\n")
		fmt.Println(err)
		return
	}
	PrintStatus("Читаю конфиг", "Done\n")

	PrintStatus("Проверяю папку \"images\"", "...")
	err = CheckDir()
	if err != nil {
		PrintStatus("Проверяю папку \"images\"", "Error\n")
		fmt.Println(err)
		return
	}
	PrintStatus("Проверяю папку \"images\"", "Done\n")

	switch rbg_json.ConfigInstance.InputFile {
	case "":
		input, err := ReadCmd()
		if err != nil {
			fmt.Println(err)
			return
		}
		PrintStatus("Удаление фона из картинки", "...")

		err = ExecCommand("python3", "rembg/remove_bg.py",
			fmt.Sprintf("./images/%s", input),
			fmt.Sprintf("%s/%s", rbg_json.ConfigInstance.OutputPath, input),
		)
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
		PrintStatus("Удаление фона из картинок", "...")

		for i, file := range input {
			fileName := strings.TrimSpace(file)

			err = ExecCommand(
				"python3", "rembg/remove_bg.py",
				fmt.Sprintf("./images/%s", fileName),
				fmt.Sprintf("%s/%s", rbg_json.ConfigInstance.OutputPath, fileName),
			)
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

func commandExist(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func PrintStatus(task string, status string) {
	fmt.Printf("\r%-50s | %-10s", task, status)
}
