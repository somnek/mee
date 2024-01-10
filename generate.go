package mee

import (
	_ "embed"
	"fmt"
	"math/rand"
	"time"
)

// Generate generates a random name by combining a random adjective and noun with the specified delimiter.
// The generated name is returned as a string.
func Generate(delimiter string) string {
	noun, adj := GetOne("noun"), GetOne("adj")
	name := fmt.Sprintf("%s%s%s", noun, delimiter, adj)
	return name
}

// GetOne returns a random word of the specified type. (noun, adj, emoji)
func GetOne(wordType string) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	var word string

	switch wordType {
	case "noun":
		return NL[rng.Intn(len(NL))]
	case "adj":
		return AL[rng.Intn(len(AL))]
	case "emoji":
		return EL[rng.Intn(len(EL))]
	}

	return word
}
