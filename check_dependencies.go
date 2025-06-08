package main

import (
	"errors"
	"fmt"
	"strings"
)

func CheckDependencies(uos, dist string) error {

	if !commandExist("python") && !commandExist("python3") {
		PrintStatus("Проверяю установленны ли необходимые файлы", "Error\n")
		fmt.Println("Необходимо установить python")
		switch uos {
		case "windows":
			fmt.Println("Установить можно по ссылке: https://www.python.org/downloads/windows/ или через Microsoft Store")
		case "darwin":
			fmt.Println("Установить можно по ссылке: https://www.python.org/downloads/macos/")
		case "linux":
			switch dist {
			case "ubuntu", "debian":
				fmt.Println(" - sudo apt update")
				fmt.Println(" - sudo apt install python3 python3-pip")
			case "arch":
				fmt.Println(" - sudo pacman -S python python-pip")
			case "fedora":
				fmt.Println(" - sudo dnf install python3 python3-pip")
			default:
				fmt.Println("Пожалуйста, установите Python вручную для вашего дистрибутива.")
			}
		}
		return errors.New("python missing")
	}
	PrintStatus("Проверяю установленны ли необходимые файлы", "Done\n")

	if !commandExist("pip") && !commandExist("pip3") {
		PrintStatus("Проверяю наличие pip", "Error\n")
		fmt.Println("Отсутствует пакетный менеджер pip")
		switch uos {
		case "windows", "darwin":
			fmt.Println("Попробуйте переустановить python")
			return errors.New("pip missing")
		case "linux":
			fmt.Print("Пробую установить pip... ")
			var err error
			switch dist {
			case "ubuntu", "debian":
				err = ExecCommand("sudo", "apt", "update", "-y")
				if err == nil {
					err = ExecCommand("sudo", "apt", "install", "-y", "python3-pip")
				}
			case "fedora":
				err = ExecCommand("sudo", "dnf", "install", "-y", "python3-pip")
			case "arch":
				err = ExecCommand("sudo", "pacman", "-S", "--noconfirm", "python-pip")
			default:
				PrintStatus("Пробую установить pip", "Error\n")
				fmt.Println("Не удалось определить дистрибутив системы, установите pip/pip3 вручную")
				return errors.New("pip missing")
			}
			if err != nil {
				PrintStatus("Пробую установить pip", "Error\n")
				return err
			}
			PrintStatus("Пробую установить pip", "Done\n")
		}
	} else {
		PrintStatus("Проверяю наличие pip", "Done\n")
	}

	checkAndInstallPackage := func(pipCmd, pkg string) error {
		task := fmt.Sprintf("Проверяю библиотеку %s", pkg)
		fmt.Printf("\r%-50s | ...       ", task)

		msg, stderrMsg, err := PipCheck(pipCmd, "show", pkg)
		if err != nil {
			fmt.Printf("Ошибка при проверке пакета %s: %v\n", pkg, err)
			return err
		}

		if strings.Contains(stderrMsg, "WARNING: Package(s) not found") || strings.Contains(msg, "not found:") {
			fmt.Printf("\r%-50s | Error     \n", task)
			installTask := fmt.Sprintf("Пробую установить %s", pkg)
			fmt.Printf("%-50s ... ", installTask)
			err = ExecCommand(pipCmd, "install", pkg)
			if err != nil {
				fmt.Printf("\r%-50s | Error     \n", installTask)
				return err
			}
			fmt.Printf("\r%-50s | Done      \n", installTask)
		} else {
			fmt.Printf("\r%-50s | Done      \n", task)
		}
		return nil
	}

	var pipCmd string
	switch uos {
	case "windows":
		pipCmd = "pip"
	case "darwin", "linux":
		pipCmd = "pip3"
	default:
		pipCmd = "pip3"
	}

	if err := checkAndInstallPackage(pipCmd, "rembg"); err != nil {
		return err
	}

	if err := checkAndInstallPackage(pipCmd, "onnxruntime"); err != nil {
		return err
	}

	return nil
}
