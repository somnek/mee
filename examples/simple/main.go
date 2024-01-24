package main

import (
	"fmt"

	"github.com/somnek/mee"
)

func main() {
	name := mee.GenerateWithDelimiter("-")
	fmt.Println(name)
}
