package util

import (
	"math/rand"
	"strings"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	r := rand.New(rand.NewSource(rand.Int63()))
	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	r := rand.New(rand.NewSource(rand.Int63()))
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "VND", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
