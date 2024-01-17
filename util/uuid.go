package util

import (
	"log"

	"github.com/google/uuid"
)

func MakeUUID() uuid.UUID {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
	}
	return uuid
}
