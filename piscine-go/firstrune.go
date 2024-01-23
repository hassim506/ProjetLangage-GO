package piscine

func FirstRune(s string) rune {
	/*val := []rune(s)
	return val[0]*/
	var first rune
	for _, c := range s {
		first = c
		break

	}
	return first
}
