package utils

import "github.com/google/uuid"

func StringToUUID(s string) uuid.UUID {
	return uuid.Must(uuid.Parse(s))
}
