package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for combination := 12; combination <= 789; combination++ {
		u := combination % 10
		d := combination / 10 % 10
		c := combination / 100 % 10

		if u > d && d > c {
			z01.PrintRune(rune(c + 48))
			z01.PrintRune(rune(d + 48))
			z01.PrintRune(rune(u + 48))

			if combination != 789 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
