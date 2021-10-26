package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			dig := 0
			for i := '1'; i <= char; i++ {
				dig++
			}
			x = x*10 + dig
		} else {
			x = 0
			return x
		}
	}
	return x
}
