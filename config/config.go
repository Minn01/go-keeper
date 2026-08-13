package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Downloads                 map[string]*int `json:"downloads"`
	ScreenshotsExpirationDays int             `json:"screenshotsExpirationDays"`
}

func DefaultConfig() Config {
	installers := 30
	media := 60
	codes := 90
	documents := 30
	others := 30

	return Config{
		Downloads: map[string]*int{
			"Code Related":      &codes,
			"Documents Related": &documents,
			"Installers":        &installers,
			"Media Related":     &media,
			"Others":            &others,
		},
		ScreenshotsExpirationDays: 30,
	}
}

func GetPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "go-keeper", "config.json"), nil
}

func Load() (Config, error) {
	// get config file path
	path, err := GetPath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)

	if os.IsNotExist(err) {
		cfg := DefaultConfig()

		if err := Save(cfg); err != nil {
			return Config{}, err
		}

		return cfg, nil
	}

	if err != nil {
		return Config{}, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("invalid config file: %w", err)
	}

	return cfg, nil
}

func Save(cfg Config) error {
	// get config file page
	path, err := GetPath()
	if err != nil {
		return err
	}

	// create the folder and file
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func PrintConfig(cfg Config) {
	fmt.Println("--- Downloads Expiration (Days) ---")
	
	// Loop over the map: key is a string, value is a pointer (*int)
	for category, daysPtr := range cfg.Downloads {
		if daysPtr != nil {
			// *daysPtr dereferences the pointer to get the actual int value
			fmt.Printf("%-20s: %d days\n", category, *daysPtr)
		} else {
			fmt.Printf("%-20s: <nil>\n", category)
		}
	}

	fmt.Println("----------------------------------")
	fmt.Printf("Screenshots Expiration: %d days\n", cfg.ScreenshotsExpirationDays)
}

func PrintConfigFilePath() {
	configFilePath, e := GetPath()

	if e != nil {
		fmt.Println("Failed to get config file path")
	}

	fmt.Println("Config: " + configFilePath)
}