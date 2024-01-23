package piscine

func IterativeFactorial(nb int) int {
	s := 1
	if nb < 0 || nb > 100 {
		return 0
	}
	if nb == 1 {
		return 1
	}

	for i := 1; i <= nb; i++ {
		s = s * i
	}
	return s
}
