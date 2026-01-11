package config

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
	FileName string `json:"file_name"`
	Level    string `json:"level"`
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
