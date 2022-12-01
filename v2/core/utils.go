package core

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	randSource := rand.New(rand.NewSource(time.Now().Unix()))
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	randBytes := make([]byte, n)
	for i, cache, remain := n-1, randSource.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			randBytes[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(randBytes)
}
