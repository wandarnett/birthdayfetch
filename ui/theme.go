package ui

import (
	"os"
)

// LoadLogo читает ascii-арт из themes/assets/<filename>.txt
func LoadLogo(filename string) (string, error) {
	data, err := os.ReadFile("themes/assets/" + filename + ".txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
