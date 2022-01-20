package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphnum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

var userName string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUser() string {
	userName = RandomString(6)
	return userName
}

func RandomEmail() string {
	return userName + "@gmail.com"
}

func RandomFirstName() string {
	return RandomString(6)
}

func RandomLastName() string {
	return RandomString(6)
}

func RandomUUID(n int) string {
	var sb strings.Builder
	k := len(alphnum)

	for i := 0; i < n; i++ {
		c := alphnum[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomLink(username string) string {
	var sb strings.Builder
	sb.WriteString("/" + username + "/")
	sb.WriteString(RandomUUID(10))

	return sb.String()
}
