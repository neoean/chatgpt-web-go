package password

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	pwd := "admin123"
	hash := Hash(pwd)
	fmt.Println(hash)
}
