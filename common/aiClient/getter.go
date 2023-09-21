package aiClient

import (
	"math/rand"
)

func GetGptClient() *GptClient {
	if len(GptClients) < 1 {
		return nil
	}

	keyCount := len(GptClients)

	index := rand.Intn(keyCount)

	return GptClients[index]
}
