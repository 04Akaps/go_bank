package util

import (
	"math/rand"
	"time"
	"strings"
)

const randomStd = "aklickdklsdoiwekfljdjls"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomint generate a random integer between min and max
func RandomInt(min,max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// randomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(randomStd)

	for i:= 0; i< n; i++ {
		c:= randomStd[rand.Intn(k)]
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
	currencies := []string{"EUR", "USD", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}