package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DatabaseUrl     string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homedir, configFileName), nil
}

func Read() Config {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		log.Fatalf("failed to get config file path: %v", err)
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("failed to decode JSON: %v", err)
	}

	return cfg
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	if err := c.write(); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}
	return nil
}

func (c *Config) write() error {
	configFileName, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
