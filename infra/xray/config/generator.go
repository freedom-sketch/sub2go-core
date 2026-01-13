package config

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/freedom-sketch/sub2go-core/dto"
	"github.com/freedom-sketch/sub2go-core/infra/database/models"
)

//go:embed templates/*.tmpl
var templates embed.FS

func prepareInboundData(inbound *models.Inbound) (dto.InboundTemplateData, error) {
	var serverNames []string
	err := json.Unmarshal([]byte(inbound.SNI), &serverNames)
	if err != nil {
		return dto.InboundTemplateData{}, err
	}
	serverNamesJSON, _ := json.Marshal(serverNames)

	var shortIds []string
	err = json.Unmarshal([]byte(inbound.ShortIds), &shortIds)
	if err != nil {
		return dto.InboundTemplateData{}, err
	}
	shortIdsJSON, _ := json.Marshal(shortIds)

	return dto.InboundTemplateData{
		Tag:             inbound.Tag,
		Port:            inbound.Port,
		Target:          inbound.Target,
		PrivateKey:      inbound.PrivateKey,
		ServerNamesJSON: string(serverNamesJSON),
		ShortIdsJSON:    string(shortIdsJSON),
	}, nil
}

func GenerateInbounConfig(inbound *models.Inbound) ([]byte, error) {
	prt, net := inbound.Protocol, inbound.Network
	tmplPath := fmt.Sprintf("templates/inbound_%s_%s_reality.tmpl", prt, net)

	tmplBytes, err := templates.ReadFile(tmplPath)
	if err != nil {
		return nil, fmt.Errorf("сould not find template: %v", err)
	}

	tmpl, err := template.New(tmplPath).Parse(string(tmplBytes))
	if err != nil {
		return nil, err
	}

	data, err := prepareInboundData(inbound)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
