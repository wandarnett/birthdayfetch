package files

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Theme  string `json:"name"`
	Lights string `json:"lights"`
}

var defaultData = Config{
	Theme:  "random",
	Lights: "individual",
}

func GetConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(configDir, "birthdayfetch", "config.json")

	if _, err := os.Stat(path); err != nil {
		return defaultData
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		log.Fatal(err)
	}

	GetThemes()

	return config
}
