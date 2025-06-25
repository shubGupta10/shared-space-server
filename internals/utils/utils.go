package utils

import (
	"github.com/google/uuid"
)

func ConvertToUUID(variable string) uuid.UUID {
	data, _ := uuid.Parse(variable)
	return data
}
