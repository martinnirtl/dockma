package emojis

import (
	"math/rand"
	"time"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}
