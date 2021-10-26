package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func talivre(queennumber, rowposition int) bool {
	for i := 0; i < queennumber; i++ {
		other_row_pos := position[i]

		if other_row_pos == rowposition ||
			other_row_pos == rowposition-(queennumber-i) ||
			other_row_pos == rowposition+(queennumber-i) {
			return false
		}
	}
	return true
}

func resolverpuzzle(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			z01.PrintRune(rune(position[i] + 48 + 1))
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if talivre(k, i) {
				position[k] = i

				resolverpuzzle(k + 1)
			}
		}
	}
}

func EightQueens() {
	resolverpuzzle(0)
}

// var (
// 	ans  [9]rune
// 	ban  [9]bool
// 	ans2 [9]int
// )

// func EightQueens() {
// 	ok := true
// 	cnt := 0
// 	for i := 1; i <= 8; i++ {
// 		if ban[i] == false {
// 			ok = false
// 		} else {
// 			cnt++
// 		}
// 	}

// 	if ok == true {
// 		for _, c := range ans {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')

// 		return
// 	}
// 	for i := '1'; i <= '8'; i++ {
// 		cur := 0
// 		for j := '1'; j <= i; j++ {
// 			cur++
// 		}
// 		if ban[cur] == false {
// 			put := true
// 			for j := 1; j <= cnt; j++ {
// 				if cur == ans2[j]-(cnt+1-j) || cur == ans2[j]+(cnt+1-j) {
// 					put = false
// 					break
// 				}
// 			}
// 			if put == true {
// 				ban[cur] = true
// 				ans[cnt+1] = i
// 				ans2[cnt+1] = cur
// 				EightQueens()
// 				ban[cur] = false
// 			}
// 		}
// 	}
// }
