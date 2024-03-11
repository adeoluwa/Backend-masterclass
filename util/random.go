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

// RandomInt generate a random Integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max -min + 1) // min->max-min
}

// The RandomString function generates a random string of specified length using characters from a
// predefined alphabet.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i :=0; i < n; i++ {
	c := alphabet[rand.Intn(k)]
	sb.WriteByte(c)
	}

	return sb.String()
}

//RandomOwner generate a random account owner 
func RandomOwner() string{
	return RandomString(6)
}

//RandomMoney generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCutrency generates a random curency code
func RandomCurrency() string {
	currencies := []string{"EUR","USD","CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}