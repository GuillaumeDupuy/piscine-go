package main

import "github.com/01-edu/z01"

func main() {
	for char := '0'; char <= '9'; char++ {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
