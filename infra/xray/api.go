package xray

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/freedom-sketch/sub2go-core/infra/xray/config"
)

func AddInbound(inbound *models.Inbound, apiAddr string) error {
	configBytes, err := config.GenerateInbounConfig(inbound)
	if err != nil {
		return fmt.Errorf("failed to generate inbound config: %w", err)
	}

	cmd := exec.Command("xray", "api", "adi", "--server="+apiAddr)

	cmd.Stdin = bytes.NewReader(configBytes)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error when executing command: %w", err)
	}

	return nil
}
