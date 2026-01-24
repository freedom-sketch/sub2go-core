package utils

import (
	"fmt"

	"github.com/freedom-sketch/sub2go-core/infra/database/models"
)

func TrafficFormat(sub *models.Subscription) string {
	total := sub.TotalTraffic
	used := sub.UsedTraffic

	if sub.TotalTraffic == 0 {
		return fmt.Sprintf("%2.f GiB / ♾️", bytesToGiB(used))
	} else {
		return fmt.Sprintf("%.2f GiB / %.2f GiB", bytesToGiB(used), bytesToGiB(total))
	}
}

func bytesToGiB(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}
