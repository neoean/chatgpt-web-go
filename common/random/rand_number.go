package random

import (
	"math/rand"
	"time"
)

const (
	randomRound = "0123456789"
)

func GenSmsCode() string {
	rand.Seed(time.Now().UnixNano())
	var slicePol []byte
	for i := 0; i < 4; i++ {
		n := rand.Intn(len(randomRound))
		slicePol = append(slicePol, randomRound[n])
	}
	return string(slicePol)
}
