package config

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/freedom-sketch/sub2go/internal/database/models"
)

func TestConfigInboundVlessRealityXHTTP(t *testing.T) {
	inbound := models.Inbound{
		ID:          1,
		ServerID:    1,
		Tag:         "test",
		Protocol:    "vless",
		Port:        443,
		Network:     "xhttp",
		Security:    "reality",
		Flow:        "",
		ShortIds:    `["1234", "a223"]`,
		PublicKey:   "dkedwedjwhfhwfjhkgwfgyu2f23fr76f",
		PrivateKey:  "de2dy2gdf26g73d674gf43g36y7feduh",
		Target:      "github.com:443",
		SNI:         `["github.com", "www.github.com"]`,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: "none",
	}

	inbConfig, err := GenerateInbounConfig(&inbound)
	if err != nil {
		t.Fatal(err)
	}

	pretty, err := json.MarshalIndent(json.RawMessage(inbConfig), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(pretty))
}

func TestConfigInboundVlessRealityTCP(t *testing.T) {
	inbound := models.Inbound{
		ID:          1,
		ServerID:    1,
		Tag:         "test",
		Protocol:    "vless",
		Port:        443,
		Network:     "tcp",
		Security:    "reality",
		Flow:        "",
		ShortIds:    `["1234", "a223"]`,
		PublicKey:   "dkedwedjwhfhwfjhkgwfgyu2f23fr76f",
		PrivateKey:  "de2dy2gdf26g73d674gf43g36y7feduh",
		Target:      "github.com:443",
		SNI:         `["github.com", "www.github.com"]`,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: "none",
	}

	inbConfig, err := GenerateInbounConfig(&inbound)
	if err != nil {
		t.Fatal(err)
	}

	pretty, err := json.MarshalIndent(json.RawMessage(inbConfig), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(pretty))
}
