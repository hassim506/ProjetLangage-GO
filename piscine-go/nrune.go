package piscine

func NRune(s string, n int) rune {
	m := []rune(s)

	for i := 0; i < len(s); i++ {
		if i+1 == n {
			return m[i]
		}
	}
	return 0
}
