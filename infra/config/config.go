package config

import (
	"encoding/json"
	"os"
)

type App struct {
	Host                  string `json:"host"`
	WebPath               string `json:"web-path"`
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SupportURL            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
	AnnounceURL           string `json:"announce-url"`
}

type XrayAPI struct {
	Port int `json:"port"`
}

type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

type TelegramBot struct {
	Token       string   `json:"token"`
	Channel     string   `json:"channel"`
	Support     string   `json:"support"`
	TgProxyURLs []string `json:"tg-proxy-urls"`
}

type Config struct {
	App         App         `json:"app"`
	XrayAPI     XrayAPI     `json:"xray-api"`
	DataBase    DataBase    `json:"database"`
	TelegramBot TelegramBot `json:"telegram-bot"`
}

var cfg *Config

// loads the config located at path into the cfg variable
func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}

	return nil
}

// returns a pointer to the config
func Get() *Config {
	return cfg
}
