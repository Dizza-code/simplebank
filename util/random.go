package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// write a function to genearate a random interger. This function takes two int64 numbers
// max and min as input and it returns random numbers
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

// function to generate randome strrings of n characters, for this we will generate alphabets
// that contains all supported characters
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

// Random Money generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 100)
}

// Random Currency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
