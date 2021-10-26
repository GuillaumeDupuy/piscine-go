package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	var count int
	for i := range os.Args {
		count = i
	}
	if count != 0 {
		os.Args = os.Args[1:]
		for i := range os.Args {
			str, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			} else {
				printStr(string(str))
			}
		}
	} else {
		str, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			printStr("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		} else {
			printStr(string(str))
		}
	}
}
