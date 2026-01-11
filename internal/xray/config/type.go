package config

type InboundTemplateData struct {
	Tag             string
	Port            int
	Target          string
	PrivateKey      string
	ServerNamesJSON string
	ShortIdsJSON    string
}
