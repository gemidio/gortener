package utils

import (
	"github.com/google/uuid"
)

type Hash struct{}

func (h *Hash) Random() string {
	return uuid.NewString()
}
