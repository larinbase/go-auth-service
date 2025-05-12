package utils

import (
	"math/rand/v2"
)

func GenerateEmailCode() int {
	return rand.IntN(899999) + 100000
}
