// This package is intended for loading the configuration
// into memory and subsequent use
package config

import (
	"encoding/json"
	"os"
)

// configuration type for the subscription service
type App struct {
	Host    string `json:"host"`     // address (domain)
	WebPath string `json:"web-path"` // path to endpoint

	// you can read more about the parameters below here -> https://www.happ.su/main/ru/dev-docs/app-management
	ProfileUpdateInterval int    `json:"profile-update-interval"` // subscription auto-renewal interval
	ProfileTitle          string `json:"profile-title"`           // subscription title
	SupportURL            string `json:"support-url"`             // support link
	ProfileWebPageUrl     string `json:"profile-web-page-url"`    // link to the main site
	Announce              string `json:"announce"`                // some kind of announcement
	AnnounceURL           string `json:"announce-url"`            // link to announcement
}

// configuration type for working with the xray API
type XrayAPI struct {
	Port int `json:"port"` // xray api port (usually 8080)
}

// configuration type for connecting to the database
type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"` // database name
}

// configuration type for the bot
type TelegramBot struct {
	Token       string   `json:"token"`
	Channel     string   `json:"channel"`       // link to the information channel
	Support     string   `json:"support"`       // support link
	TgProxyURLs []string `json:"tg-proxy-urls"` // list of URLs for connecting to a proxy in Telegram
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
