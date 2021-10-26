package piscine

func Split(s, sep string) []string {
	n := 0
	m := 0
	for i := range s {
		n = i + 1
	}
	for i := range sep {
		m = i + 1
	}
	foundSep := true
	for foundSep {
		foundSep = false
		for i := 0; i < n-m; i++ {
			if s[i:i+m] == sep {
				foundSep = true
				s = s[:i] + " " + s[i+m:]
				n -= m
				break
			}
		}
	}

	return SplitWhiteSpaces(s)
}
