package owlgram_server

import (
	"math/rand"
	"time"
)

func RandName() string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("0123456789abcdefghijkl0123456789mnopqrs0123456789tuvwxyz0123456789")
	b := make([]rune, 18)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
