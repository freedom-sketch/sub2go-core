package utils

import (
	"strconv"

	"github.com/google/uuid"
)

func IntToUUID(n int64) uuid.UUID {
	namespace := uuid.NameSpaceURL
	input := strconv.FormatInt(n, 10)
	return uuid.NewSHA1(namespace, []byte(input))
}
