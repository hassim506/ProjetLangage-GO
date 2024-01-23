package piscine

func StrRev(s string) string {
	n := len(s)
	words := make([]rune, n)
	for _, word := range s {
		n--
		words[n] = word
	}
	return string(words[n:])
}
