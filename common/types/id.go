package types

import (
	"math/rand"
	"time"
)

func GenerateUID() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 100000 // 最小值（6位数）
	max := 999999 // 最大值（6位数）
	return int64(rand.Intn(max-min+1) + min)
}
