package utils

import (
	"fmt"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/google/uuid"
)

func GenerateSubscriptionKey(UserUUID uuid.UUID) string {
	cfg := config.Get()

	return fmt.Sprintf("https://%s/%s/%s", cfg.App.Host, cfg.App.WebPath, UserUUID.String())
}
