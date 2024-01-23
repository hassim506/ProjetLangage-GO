package piscine

func LastRune(s string) rune {
	val := []rune(s)
	return val[len(val)-1]
}
