package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0])[2:]
	for _, letter := range runes {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
