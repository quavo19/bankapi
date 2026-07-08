package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of given length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}


// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomCurrency generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GHS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}


// RandomeInt generates a random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}