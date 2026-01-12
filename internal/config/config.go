package config

import (
	"encoding/json"
	"os"
)

type Subscription struct {
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SupportURL            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
	AnnounceURL           string `json:"announce-url"`
}

type API struct {
	Host    string `json:"host"`
	WebPath string `json:"web_path"`
}

type Logging struct {
	Level string `json:"level"`
	Path  string `json:"path"`
}

type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

type Config struct {
	API          API          `json:"api"`
	Subscription Subscription `json:"subscription"`
	Logging      Logging      `json:"logging"`
	DataBase     DataBase     `json:"data-base"`
}

// Loads data from the configuration into a structure and returns a pointer to it
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
