package util

import (
	"fmt"

	"github.com/google/uuid"
)

func MakeUUID() uuid.UUID {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	return uuid
}
