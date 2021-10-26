package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		PrintStr(s)
		z01.PrintRune('\n')
	}
}

// func PrintWordsTables(a []string) {
// 	for _, s := range a {
// 		for _, r := range []rune(s) {
// 			z01.PrintRune(r)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
