package utils

import (
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/google/uuid"
)

func GenerateSubscriptionKey(UserUUID uuid.UUID) string {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Panicf("Failed to load config: %v", err)
	}

	return fmt.Sprintf("https://%s/%s/%s", cfg.API.Host, cfg.API.WebPath, UserUUID.String())
}
