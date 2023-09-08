package gpt

import (
	"math/rand"
)

func GetClient() *GptClient {
	if len(Clients) < 1 {
		return nil
	}

	keyCount := len(Clients)

	index := rand.Intn(keyCount)

	return Clients[index]
}
