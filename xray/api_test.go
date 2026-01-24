package xray

import (
	"testing"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/freedom-sketch/sub2go-core/xray/templates"
)

func TestAddInbound(t *testing.T) {
	cfg, err := config.Load("config.json")
	if err != nil {
		t.Fatal(err)
	}

	client := XrayAPI{}
	err = client.Init(cfg.XrayAPI.Port)
	defer client.Close()
	if err != nil {
		t.Fatal(err)
	}

	testInbound := models.Inbound{
		ID:          1,
		ServerID:    1,
		Tag:         "test",
		Protocol:    "vless",
		Port:        5678,
		Network:     "tcp",
		Security:    "reality",
		Flow:        "xtls-rprx-vision",
		ShortIds:    `["abcd"]`,
		PublicKey:   "hvSSKxEtssgDrXsVOUOtmGUGqK0lKRWXjBl5xxIJ0w8",
		PrivateKey:  "WFQIB2_563xyVcPICdWiPuGAKMcpv9yQZLtJbYCvdn8",
		Target:      "github.com:443",
		SNI:         `["github.com", "www.github.com"]`,
		Description: "test",
		IsActive:    true,
	}

	inboundConfig, err := templates.GenerateInboundConfig(&testInbound)
	if err != nil {
		t.Fatal(err)
	}

	err = client.AddInbound(inboundConfig)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("inbound successfully created")
}
