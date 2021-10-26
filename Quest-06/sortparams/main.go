package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	counter := 0
	for i := range os.Args {
		counter++
		_ = i
	}
	for i := 1; i < counter; i++ {
		for j := 1; j < counter-i; j++ {
			if os.Args[j] > os.Args[j+1] {
				os.Args[j], os.Args[j+1] = os.Args[j+1], os.Args[j]
			}
		}
	}
	for i := 1; i < counter; i++ {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune(10)
	}
}
