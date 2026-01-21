package config

import (
	"encoding/json"
	"os"
)

type API struct {
	Host                  string `json:"host"`
	WebPath               string `json:"web_path"`
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SupportURL            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
	AnnounceURL           string `json:"announce-url"`
}

type Logging struct {
	Level           string `json:"level"`
	SubscriptionAPI string `json:"subscription-api"`
	AgentsAPI       string `json:"agents-api"`
	DataBase        string `json:"data-base"`
	TelegramBot     string `json:"telegram-bot"`
}

type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

type TelegramBot struct {
	Token string `json:"token"`
}

type Config struct {
	API         API         `json:"api"`
	Logging     Logging     `json:"logging"`
	DataBase    DataBase    `json:"data-base"`
	TelegramBot TelegramBot `json:"telegram-bot"`
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
