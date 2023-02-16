package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate() string {
	salt := make([]rune, 12)
	for i := range salt {
		salt[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(salt)
}

func main() {
	fmt.Println(Generate())
}
