package main

import (
	"os"
)

func main() {
	os.Args = os.Args[1:]
	count := 0
	for i := range os.Args {
		count = i
	}
	if count > 1 && IsNumeric(os.Args[0]) && IsNumeric(os.Args[2]) {
		a := Atoi(os.Args[0])
		b := Atoi(os.Args[2])
		Operation(a, b, os.Args[1])
	}
}

func IsNumeric(s string) bool {
	for _, l := range s {
		if !(47 < l && l < 58) && l != '-' {
			return false
		}
	}
	return true
}

func Operation(a, b int, op string) {
	if a < 9223372036854775807 && a > -9223372036854775807 {
		if op == "/" {
			if b != 0 {
				PrintStr(intToString(a / b))
			} else {
				PrintStr("No division by 0\n")
			}
		} else if op == "%" {
			if b != 0 {
				PrintStr(intToString(a % b))
			} else {
				PrintStr("No modulo by 0\n")
			}
		} else if op == "*" {
			PrintStr(intToString(a * b))
		} else if op == "+" {
			t := a + b
			if a != t {
				PrintStr(intToString(t))
			}
		} else if op == "-" {
			t := a - b
			if a != t {
				PrintStr(intToString(t))
			}
		}
	}
}

func Atoi(s string) int {
	var res int
	y := 1
	l := StrLen(s)
	var negative bool
	if l == 0 {
		return 0
	}
	n := 0
	for s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			negative = true
			l -= 1
			s = s[1:]
			n += 1
		} else if s[0] == '+' {
			negative = false
			l -= 1
			s = s[1:]
			n += 1
		}
		if s == "" || n == 2 {
			return 0
		}
	}
	if s == "" {
		return 0
	}
	for i := l - 1; i >= 0; i-- {
		t := (int(s[i]) - 48)
		good := false
		for g := 0; g < 10; g++ {
			if g == t && !good {
				good = true
			}
		}
		if !good {
			return 0
		}
		res += (int(s[i]) - 48) * y
		y *= 10
	}
	if negative {
		res *= -1
	}
	return res
}

func StrLen(s string) int {
	var count int
	for j := range []rune(s) {
		count = j + 1
	}
	return count
}

func PrintStr(s string) {
	os.Stdout.WriteString(s)
}

func intToString(num int) string {
	var res string
	var neg bool
	if num == 0 {
		return "0\n"
	}
	if num < 0 {
		neg = true
		num *= -1
		if num < 0 {
			return ""
		}
	}
	for num > 0 {
		num2 := num % 10
		t := rune(num2 + 48)
		res = string(t) + res
		num /= 10
	}
	if neg {
		res = "-" + res
	}
	return res + "\n"
}
