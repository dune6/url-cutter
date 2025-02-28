package random

import (
	"math/rand"
	"strings"
	"time"
)

const (
	minLetter = 97  // a letter
	maxLetter = 123 // z+1 letter
)

func NewRandomString(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var str strings.Builder
	for range length {
		random := rnd.Intn(maxLetter-minLetter) + minLetter
		str.WriteRune(rune(random))
	}

	return str.String()
}
