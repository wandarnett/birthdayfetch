package themes

import (
	"fmt"
	"os"
	"strings"
)

func SummerTheme() {
	data, err := os.ReadFile("themes/assets/summer.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		var coloredLine string
		for _, ch := range line {
			switch ch {
			case '*':
				coloredLine += "\033[1;33m" + string(ch) + "\033[0m"
			case '.':
				coloredLine += "\033[1;97m" + string(ch) + "\033[0m"
			case '|', '/', '\\':
				coloredLine += "\033[1;33m" + string(ch) + "\033[0m"
			case '~', '_', '-':
				coloredLine += "\033[1;33m" + string(ch) + "\033[0m"
			default:
				coloredLine += "\033[1;33m" + string(ch) + "\033[0m"
			}
		}
		fmt.Println(coloredLine)
	}
}

func AutumnTheme() {
	data, _ := os.ReadFile("themes/assets/autumn.txt")
	lines := strings.Split(string(data), "\n")

	colors := []string{"\033[0;33m", "\033[0;31m", "\033[1;33m"}

	for i, line := range lines {
		c := colors[i%len(colors)]
		fmt.Printf("%s%s\033[0m\n", c, line)
	}
}

func SpringTheme() {
	data, err := os.ReadFile("themes/assets/spring.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	totalLines := len(lines)

	for i, line := range lines {
		var coloredLine string
		for _, ch := range line {
			switch ch {
			case '/', '\\': // верхушка цветка — розовый
				coloredLine += "\033[1;95m" + string(ch) + "\033[0m"
			case '|': // стебель — зелёный
				coloredLine += "\033[1;32m" + string(ch) + "\033[0m"
			case '\'': // цветок — темно-зелёный
				coloredLine += "\033[0;32m" + string(ch) + "\033[0m"
			case '^': // верхушка или трава — проверяем по позиции строки
				if i >= totalLines*3/4 { // нижняя часть — трава
					coloredLine += "\033[0;32m" + string(ch) + "\033[0m"
				} else { // верхушка — розовый
					coloredLine += "\033[1;95m" + string(ch) + "\033[0m"
				}
			default:
				coloredLine += string(ch) // пробелы и остальные символы
			}
		}
		fmt.Println(coloredLine)
	}
}

func WinterTheme() {
	data, _ := os.ReadFile("themes/assets/winter.txt")
	lines := strings.Split(string(data), "\n")

	colors := []string{"\033[1;36m", "\033[0;34m", "\033[1;37m"}

	for i, line := range lines {
		c := colors[i%len(colors)]
		fmt.Printf("%s%s\033[0m\n", c, line)
	}
}
