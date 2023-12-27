package mee

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate return <noun><delimiter><adjective>
func Generate(delimiter string) string {
	name := fmt.Sprintf("%s%s%s", GetNoun(), delimiter, GetAdj())
	return name
}

func GetAdj() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return AL[rng.Intn(len(AL))]
}

func GetNoun() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return NL[rng.Intn(len(NL))]
}

func OneWord() string {
	// this function get 1 random word from wordlist
	return ""
}

func OneEmoji() string {
	// this function get 1 random word from wordlist
	return ""
}
