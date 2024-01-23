package piscine

func RecursivePower(nb int, power int) int {
	a := nb
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}

	a = a * RecursivePower(nb, (power-1))

	return a
}
