package random

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
)

// seeds the pseudo random generator for uniqueness
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string for fixed length
func RandomString(sLen int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < sLen; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomAmount returns a random amount between [0-n)
func RandomAmount(n int) int {
	return rand.Intn(n)
}

// RandomOwner returns a random owner name of fixed length
func RandomOwner(n int) string {
	return RandomString(n)
}

// RandomCurrency returns a random currency from [USD, INR, EUR]
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
