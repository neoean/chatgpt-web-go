package random

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	randomRound = "0123456789"

	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
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

func GenCarmiKey() string {
	return fmt.Sprintf("CA%v%v", time.Now().Unix(), generateRandomString(6))
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
