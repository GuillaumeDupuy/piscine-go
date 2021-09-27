package piscine

func NRune(s string, n int) rune {
	res := []rune(s)
	for index, valeur := range res {
		if index == n-1 {
			return valeur
		}
	}
	return 0
}
