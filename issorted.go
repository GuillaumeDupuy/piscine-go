package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	result := 1
	result2 := 1
	result3 := 1
	for x, y := range tab {
		if x != len(tab)-1 {
			if f(y, tab[x+1]) < 0 {
				result++
			}
			if f(y, tab[x+1]) > 0 {
				result2++
			}
			if f(y, tab[x+1]) == 0 {
				result3++
			}
		}
	}
	return result == len(tab) || result2 == len(tab) || result3 == len(tab)
}
