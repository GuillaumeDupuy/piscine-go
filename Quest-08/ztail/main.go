package main

import (
	"fmt"
	"os"
)

var (
	errr  bool
	brute int
)

func printStr(s string) {
	fmt.Print(s)
}

func lenT(t []string) int {
	count := 0
	for i := range t {
		count = i
	}
	return count
}

func main() {
	os.Args = os.Args[1:]
	if os.Args[0] == "-c" {
		c := BasicAtoi(os.Args[1])
		os.Args = os.Args[2:]
		t := lenT(os.Args)
		if t == 0 {
			str, k := Read(os.Args[0])
			if k != 0 {
				if k-c+1 < 0 {
					printStr(str)
				} else {
					printStr(string([]rune(str)[k-c:]))
				}
			} else {
				fmt.Println()
			}
		} else {
			for _, i := range os.Args {
				if errr {
					fmt.Println()
				}
				str, k := Read(i)
				if k != 0 {
					if errr || brute > 0 {
						fmt.Println()
					}
					printStr("==> " + i + " <==\n")
					if k-c+1 < 0 {
						printStr(str)
					} else {
						printStr(string([]rune(str)[k-c:]))
					}
					brute++
				}
			}
		}
		if errr {
			os.Exit(1)
		}
	} else {
		t := lenT(os.Args)
		if t == 0 {
			str, _ := Read(os.Args[0])
			printStr(str)
		} else {
			for _, i := range os.Args {
				str, k := Read(i)
				if k != 0 {
					printStr("==> " + i + " <==\n")
					printStr(str)
				}
			}
		}
		if errr {
			os.Exit(1)
		}
	}
}

func Read(filename string) (string, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err.Error())
		errr = true
	} else {
		defer file.Close()
		fileinfo, err := file.Stat()
		if err != nil {
			printStr(err.Error() + "\n")
			os.Exit(1)
		}
		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)
		bytesread, err := file.Read(buffer)
		if err != nil {
			printStr(err.Error() + "\n")
			os.Exit(1)
		}
		return string(buffer), bytesread
	}
	return "", 0
}

func BasicAtoi(s string) int {
	var res int
	l := StrLen(s)
	y := 1
	for i := l - 1; i >= 0; i-- {
		res += (int(s[i]) - 48) * y
		y *= 10
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
