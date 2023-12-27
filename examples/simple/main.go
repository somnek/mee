package main

import (
	"fmt"

	"github.com/somnek/mee"
)

func main() {
	name := mee.Generate("-")
	fmt.Println(name)
}
